package rest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/sjson"
)

type InParams struct {
	Address    string `json:"address"`
	Name       string `json:"name"`
	IsRequired bool   `json:"is_required"`
	Result     string
	Type       string `json:"type"`
	Auth       string `json:"auth"`
}

type Endpoint struct {
	Name        string      `json:"name"`
	Url         string      `json:"url"`
	Method      string      `json:"method"`
	AuthService string      `json:"auth_service"`
	AuthType    string      `json:"auth_type"`
	InParams    []InParams  `json:"in_params"`
	OutParams   []OutParams `json:"out_params"`
}

type OutParams struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Scheema string `json:"scheema"`
}

func Request(serviceName string, url string, method string, inParams *[]InParams) (string, error) {

	//TODO implement GET, PUT, DELETE
	if method == "POST" {
		return post(serviceName, url, inParams)
	}

	return "", errors.New("METHOD_UNKNOWN")
}

func post(serviceName string, url string, inParams *[]InParams) (string, error) {

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

func convertToBody(inParams *[]InParams) (string, error) {

	body := "{}"
	for _, s := range *inParams {
		if s.Type == "body" {
			body, _ = sjson.Set(body, s.Address, s.Result)
		}
	}

	return body, nil
}

func addToHeader(req *http.Request, inParams *[]InParams) error {

	for _, s := range *inParams {
		if s.Type == "header" {
			req.Header.Set(s.Address, s.Result)
		}
	}

	return nil
}
