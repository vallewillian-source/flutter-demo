package rest

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

func Request(serviceName string, url string, method string, authType string, inParams *[]models.InParams) (string, error) {

	//TODO implement GET, PUT, DELETE
	if method == "POST" {
		return post(serviceName, url, authType, inParams)
	}

	return "", errors.New("METHOD_UNKNOWN")
}

func post(serviceName string, url string, authType string, inParams *[]models.InParams) (string, error) {

	// creating body data
	body, err := convertToBody(inParams)
	if err != nil {
		return "", err
	}
	requestBody := bytes.NewBuffer([]byte(body))

	// preparing a http request
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return "", err
	}

	// getting auth data
	if authType == "BEARER_TOKEN" {
		getBearerAuth(serviceName, inParams)
	} else if authType == "NONE" {
	} else {
		return "", errors.New("AUTH_TYPE_INVALID")
	}

	// adding params to header
	addToHeader(req, inParams)

	// execute a http request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(responseBody), nil
}

func getBearerAuth(serviceName string, inParams *[]models.InParams) {

	// open json file
	jsonFile, err := os.Open("./jsons/auth/" + serviceName + ".json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	// convert to struct
	var auth models.BearerLoginFile
	json.Unmarshal(byteValue, &auth)

	for i, s := range *inParams {
		if s.Auth == "auth_token" {
			(*inParams)[i].Result = auth.AuthToken
		} else if s.Auth == "auth_user_id" {
			(*inParams)[i].Result = auth.AuthUserId
		}
	}

}

func convertToBody(inParams *[]models.InParams) (string, error) {

	body := "{}"
	for _, s := range *inParams {
		if s.Type == "body" {
			body, _ = sjson.Set(body, s.Address, s.Result)
		}
	}

	return body, nil
}

func addToHeader(req *http.Request, inParams *[]models.InParams) error {

	for _, s := range *inParams {
		if s.Type == "header" {
			req.Header.Set(s.Address, s.Result)
		}
	}

	return nil
}
