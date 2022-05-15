package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/sjson"
	"github.com/vallewillian-source/go-sofa-data-studio/models"
)

func Request(service_name string, url string, method string, auth_type string, in_params *[]models.In_params) (string, error) {

	//TODO implement GET, PUT, DELETE
	if method == "POST" {
		return post(service_name, url, auth_type, in_params)
	}

	return "", errors.New("METHOD_UNKNOWN")
}

func post(service_name string, url string, auth_type string, in_params *[]models.In_params) (string, error) {

	// creating body data
	body, err := convert_to_body(in_params)
	if err != nil {
		return "", err
	}
	responseBody := bytes.NewBuffer([]byte(body))

	// preparing a http request
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, responseBody)
	if err != nil {
		return "", err
	}

	// getting auth data
	if auth_type == "BEARER_TOKEN" {
		get_bearer_auth(service_name, in_params)
	} else if auth_type == "NONE" {
	} else {
		return "", errors.New("AUTH_TYPE_INVALID")
	}

	// adding params to header
	add_to_header(req, in_params)

	// execute a http request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read the response body
	response_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(response_body), nil
}

func get_bearer_auth(service_name string, in_params *[]models.In_params) {

	// open json file
	jsonFile, err := os.Open("./jsons/auth/" + service_name + ".json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	// convert to struct
	var auth models.Bearer_login_file
	json.Unmarshal(byteValue, &auth)

	for i, s := range *in_params {
		if s.Auth == "auth_token" {
			(*in_params)[i].Result = auth.Auth_token
		} else if s.Auth == "auth_user_id" {
			(*in_params)[i].Result = auth.Auth_user_id
		}
	}

}

func convert_to_body(in_params *[]models.In_params) (string, error) {

	body := "{}"
	for _, s := range *in_params {
		if s.Type == "body" {
			body, _ = sjson.Set(body, s.Address, s.Result)
		}
	}

	return body, nil
}

func add_to_header(req *http.Request, in_params *[]models.In_params) error {

	for _, s := range *in_params {
		if s.Type == "header" {
			req.Header.Set(s.Address, s.Result)
		}
	}

	return nil
}
