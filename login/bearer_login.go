package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/vallewillian-source/go-sofa-data-studio/helpers"
	"github.com/vallewillian-source/go-sofa-data-studio/io"
	"github.com/vallewillian-source/go-sofa-data-studio/models"
)

func Bearer_login(file string) {
	print("\nlogin_bearer()")

	// open json file
	jsonFile, err := os.Open("./jsons/" + file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n..Successfully opened " + file)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	// convert to struct
	var login models.Bearer_login_service
	json.Unmarshal(byteValue, &login)

	// request params from user
	in_parameters := login.Login_endpoint.In_params
	io.Request_params(&in_parameters)

	// make a http request
	response_body, err := helpers.Request(login.Login_endpoint.Url, login.Login_endpoint.Method, &in_parameters)
	if err != nil {
		fmt.Println(err)
	}

	// getting token and user id
	auth_token := gjson.Get(response_body, login.Login_endpoint.Out_params.Auth_token)
	auth_user_id := gjson.Get(response_body, login.Login_endpoint.Out_params.Auth_user_id)

	// saving auth data to file
	tokens, _ := sjson.Set("{}", "auth_token", auth_token.String())
	tokens, _ = sjson.Set(tokens, "auth_user_id", auth_user_id.String())

	os.WriteFile("jsons/auth/"+login.Service_name+".json", []byte(tokens), 0644)

}
