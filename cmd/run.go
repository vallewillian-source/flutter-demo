package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vallewillian-source/go-sofa-data-studio/internal/auth"
	"github.com/vallewillian-source/go-sofa-data-studio/internal/io"
	"github.com/vallewillian-source/go-sofa-data-studio/internal/scheema"

	"github.com/tidwall/gjson"
	"github.com/vallewillian-source/go-sofa-data-studio/internal/rest"
)

func Run(serviceName string, endpointName string) error {

	// open api json file
	apiJsonFile, err := os.Open("./jsons/services/" + serviceName + "/auth.json")
	if err != nil {
		return err
	}

	apiByteValue, _ := ioutil.ReadAll(apiJsonFile)
	defer apiJsonFile.Close()

	// convert to struct
	var api apiFile
	json.Unmarshal(apiByteValue, &api)

	// open endpoint json file
	jsonFile, err := os.Open("./jsons/services/" + serviceName + "/endpoints/" + endpointName + ".json")
	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	// convert to struct
	var endpoint rest.Endpoint
	json.Unmarshal(byteValue, &endpoint)

	// request params from user
	inParameters := endpoint.InParams
	io.FetchParams(&inParameters)

	// getting auth data
	err = auth.FetchAuthParameters(endpoint.AuthService, api.AuthType, &inParameters)
	if err != nil {
		return err
	}

	// make a http request
	responseBody, err := rest.Request(endpoint.AuthService, endpoint.Url, endpoint.Method, &inParameters)
	if err != nil {
		return err
	}

	// display result
	show(responseBody, &endpoint.OutParams)

	return nil

}

func show(response string, outParams *[]rest.OutParams) {
	for _, param := range *outParams {
		value := gjson.Get(response, param.Address)
		if json.Valid([]byte(value.String())) {
			if len(param.Scheema) > 0 {
				// print scheema
				scheema.ShowScheema(param.Scheema, value.String())
			} else {
				// unknown scheema
				fmt.Printf("\n%s: %s \n", param.Name, value.String())
			}
		} else {
			// print basic value
			fmt.Printf("\n%s: %s \n", param.Name, value.String())
		}

	}
}
