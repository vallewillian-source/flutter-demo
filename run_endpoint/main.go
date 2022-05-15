package run_endpoint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
	"github.com/vallewillian-source/go-sofa-data-studio/helpers"
	"github.com/vallewillian-source/go-sofa-data-studio/io"
	"github.com/vallewillian-source/go-sofa-data-studio/models"
)

func Execute(file string) {
	print("\nrun_endpoint()")

	// open json file
	jsonFile, err := os.Open("./jsons/" + file)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	// convert to struct
	var endpoint models.Endpoint
	json.Unmarshal(byteValue, &endpoint)

	// request params from user
	in_parameters := endpoint.In_params
	io.Request_params(&in_parameters)

	// make a http request
	response_body, err := helpers.Request(endpoint.Auth_service, endpoint.Url, endpoint.Method, endpoint.Auth_type, &in_parameters)
	if err != nil {
		fmt.Println(err)
	}

	// display result
	show(response_body, &endpoint.Out_params)

}

func show(response string, out_params *[]models.Out_params) {
	for _, param := range *out_params {
		fmt.Printf("%s: %s \n", param.Name, gjson.Get(response, param.Address))
	}
}
