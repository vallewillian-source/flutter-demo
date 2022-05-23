package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/vallewillian-source/go-sofa-data-studio/internal/auth"
)

func Login(serviceName string) error {

	// open json file
	jsonFile, err := os.Open("./jsons/services/" + serviceName + "/auth.json")
	if err != nil {
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	// convert to struct
	var api apiFile
	json.Unmarshal(byteValue, &api)

	err = auth.Login(serviceName, api.AuthType)
	if err != nil {
		return err
	}

	return nil
}
