package helpers

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/sjson"
	"github.com/vallewillian-source/go-sofa-data-studio/models"
)

func Request(url string, method string, in_params *[]models.In_params) (string, error) {
	if method == "POST" {
		return post(url, in_params)
	}

	return "", errors.New("METHOD_UNKNOWN")
}

func post(url string, in_params *[]models.In_params) (string, error) {

	// creating body data
	body, err := convert_to_body(in_params)
	if err != nil {
		return "", err
	}

	// making a http request
	responseBody := bytes.NewBuffer([]byte(body))

	resp, err := http.Post(url, "application/json", responseBody)
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

func convert_to_body(in_params *[]models.In_params) (string, error) {

	body := "{}"
	for _, s := range *in_params {
		body, _ = sjson.Set(body, s.Address, s.Result)
	}

	return body, nil
}
