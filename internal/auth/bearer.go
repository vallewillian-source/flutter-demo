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
)

func bearerLogin(serviceName string) {

	// open json file
	jsonFile, err := os.Open("./jsons/services/" + serviceName + "/auth.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	// convert to struct
	var login bearerLoginService
	json.Unmarshal(byteValue, &login)

	// request params from user
	inParameters := login.LoginEndpoint.InParams
	io.FetchParams(&inParameters)

	// make a http request
	responseBody, err := rest.Request(login.ServiceName, login.LoginEndpoint.Url, login.LoginEndpoint.Method, &inParameters)
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

func bearerFetchAuthParameters(serviceName string, inParams *[]rest.InParams) {

	// open json file
	jsonFile, err := os.Open("./jsons/auth/" + serviceName + ".json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	// convert to struct
	var auth bearerLoginFile
	json.Unmarshal(byteValue, &auth)

	for i, s := range *inParams {
		// TODO get parameters from auth scheema
		if s.Auth == "auth_token" {
			(*inParams)[i].Result = auth.AuthToken
		} else if s.Auth == "auth_user_id" {
			(*inParams)[i].Result = auth.AuthUserId
		}
	}

}
