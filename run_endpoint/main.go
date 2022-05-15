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
		value := gjson.Get(response, param.Address)
		if json.Valid([]byte(value.String())) {
			if len(param.Scheema) > 0 {
				// print scheema
				show_scheema(param.Scheema, value.String())
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

//TODO implement array of scheemas
func show_scheema(scheema_name string, value string) {

	// open json file
	// TODO implement subfolders for jsons
	// TODO implement cache
	jsonFile, err := os.Open("./jsons/" + scheema_name + ".json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	// convert to struct
	var scheema models.Scheema
	json.Unmarshal(byteValue, &scheema)

	// printing title
	fmt.Printf("\n\n[%s]", scheema.Name)

	// printing fields
	for _, field := range scheema.Fields {

		field_value := gjson.Get(value, field.Address)
		if json.Valid([]byte(field_value.String())) {
			if len(field.Scheema) > 0 {
				// print scheema
				show_scheema(field.Scheema, field_value.String())
			} else {
				// unknown scheema
				fmt.Printf("\n%s: %s", field.Name, field_value.String())
			}
		} else {
			// print basic value
			fmt.Printf("\n%s: %s", field.Name, field_value.String())
		}

	}
	fmt.Printf("\n")

}
