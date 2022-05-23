package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/vallewillian-source/go-sofa-data-studio/internal/io"
	"github.com/vallewillian-source/go-sofa-data-studio/internal/rest"
	"github.com/vallewillian-source/go-sofa-data-studio/models"
)

func BearerLogin(file string) {
	print("\nlogin_bearer()")

	// open json file
	jsonFile, err := os.Open("./jsons/" + file)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	// convert to struct
	var login models.BearerLoginService
	json.Unmarshal(byteValue, &login)

	// request params from user
	inParameters := login.LoginEndpoint.InParams
	io.FetchParams(&inParameters)

	// make a http request
	responseBody, err := rest.Request(login.ServiceName, login.LoginEndpoint.Url, login.LoginEndpoint.Method, "NONE", &inParameters)
	if err != nil {
		fmt.Println(err)
	}

	// getting token and user id
	authToken := gjson.Get(responseBody, login.LoginEndpoint.OutParams.AuthToken)
	authUserId := gjson.Get(responseBody, login.LoginEndpoint.OutParams.AuthUserId)

	// saving auth data to file
	tokens, _ := sjson.Set("{}", "auth_token", authToken.String())
	tokens, _ = sjson.Set(tokens, "auth_user_id", authUserId.String())

	os.WriteFile("jsons/auth/"+login.ServiceName+".json", []byte(tokens), 0644)

}
