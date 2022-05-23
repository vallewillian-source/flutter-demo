package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vallewillian-source/go-sofa-data-studio/internal/io"

	"github.com/tidwall/gjson"
	"github.com/vallewillian-source/go-sofa-data-studio/internal/rest"
	"github.com/vallewillian-source/go-sofa-data-studio/models"
)

func Run(file string) {
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
	inParameters := endpoint.InParams
	io.FetchParams(&inParameters)

	// make a http request
	responseBody, err := rest.Request(endpoint.AuthService, endpoint.Url, endpoint.Method, endpoint.AuthType, &inParameters)
	if err != nil {
		fmt.Println(err)
	}

	// display result
	show(responseBody, &endpoint.OutParams)

}

func show(response string, outParams *[]models.OutParams) {
	for _, param := range *outParams {
		value := gjson.Get(response, param.Address)
		if json.Valid([]byte(value.String())) {
			if len(param.Scheema) > 0 {
				// print scheema
				showScheema(param.Scheema, value.String())
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
func showScheema(scheemaName string, value string) {

	// open json file
	// TODO implement cache
	jsonFile, err := os.Open("./jsons/" + scheemaName + ".json")
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

		fieldValue := gjson.Get(value, field.Address)
		if json.Valid([]byte(fieldValue.String())) {
			if len(field.Scheema) > 0 {
				// print scheema
				showScheema(field.Scheema, fieldValue.String())
			} else {
				// unknown scheema
				fmt.Printf("\n%s: %s", field.Name, fieldValue.String())
			}
		} else {
			// print basic value
			fmt.Printf("\n%s: %s", field.Name, fieldValue.String())
		}

	}
	fmt.Printf("\n")

}
