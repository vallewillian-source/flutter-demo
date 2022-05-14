package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vallewillian-source/go-sofa-data-studio/io"
)

func Bearer_login(file string) {
	print("\nlogin_bearer()")

	jsonFile, err := os.Open("./jsons/" + file)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n..Successfully opened " + file)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var login bearer_login_service
	json.Unmarshal(byteValue, &login)

	in_parameters := login.Login_endpoint.In_params

	fmt.Println("\n...Result: " + string(in_parameters[0].Name))

	io.Request_params(in_parameters)

	defer jsonFile.Close()

}
