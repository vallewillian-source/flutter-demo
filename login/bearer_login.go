package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vallewillian-source/go-sofa-data-studio/io"
	"github.com/vallewillian-source/go-sofa-data-studio/models"
)

func Bearer_login(file string) {
	print("\nlogin_bearer()")

	jsonFile, err := os.Open("./jsons/" + file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n..Successfully opened " + file)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	var login models.Bearer_login_service
	json.Unmarshal(byteValue, &login)

	in_parameters := login.Login_endpoint.In_params
	params := io.Request_params(in_parameters)

	for _, s := range *params {
		fmt.Println(s.Name, s.Result)
	}

}
