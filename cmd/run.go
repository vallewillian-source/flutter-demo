package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/vallewillian-source/go-sofa-data-studio/internal/auth"
	"github.com/vallewillian-source/go-sofa-data-studio/internal/io"
	"github.com/vallewillian-source/go-sofa-data-studio/internal/scheema"

	"github.com/tidwall/gjson"
	"github.com/vallewillian-source/go-sofa-data-studio/internal/rest"
)

func Run(serviceName string, endpointName string) error {

	// open api json file
	apiJsonFile, err := os.Open("./jsons/services/" + serviceName + "/api.json")
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

	// generate result
	result, err := generateResult(serviceName, responseBody, &endpoint.OutParams)
	if err != nil {
		return err
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return err
	}

	os.WriteFile("jsons/results/"+serviceName+"_"+endpointName+"_"+strconv.Itoa(int(time.Now().Unix()))+".json", []byte(resultJson), 0644)

	return nil

}

func generateResult(serviceName string, response string, outParams *[]rest.OutParams) (map[string]interface{}, error) {
	var result map[string]interface{} = make(map[string]interface{})

	for _, param := range *outParams {
		value := gjson.Get(response, param.Address)
		if json.Valid([]byte(value.String())) {
			if len(param.Scheema) > 0 {
				// print scheema
				var err error
				result[param.Name], err = scheema.GenerateScheema(serviceName, param.Scheema, value.String())
				if err != nil {
					result[param.Name] = "Err"
				}
			} else {
				// unknown scheema
				result[param.Name] = value.String()
			}
		} else {
			// print basic value
			result[param.Name] = value.String()
		}

	}

	return result, nil
}
