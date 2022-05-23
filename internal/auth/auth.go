package auth

import (
	"errors"

	"github.com/vallewillian-source/go-sofa-data-studio/internal/rest"
)

type bearerLoginService struct {
	ServiceName   string              `json:"service_name"`
	AuthType      string              `json:"auth_type"`
	LoginEndpoint bearerLoginEndpoint `json:"login_endpoint"`
}

type bearerLoginEndpoint struct {
	Url       string               `json:"url"`
	Method    string               `json:"method"`
	InParams  []rest.InParams      `json:"in_params"`
	OutParams bearerLoginOutParams `json:"out_params"`
}

type bearerLoginOutParams struct {
	AuthToken  string `json:"auth_token"`
	AuthUserId string `json:"auth_user_id"`
}

type bearerLoginFile struct {
	AuthToken  string `json:"auth_token"`
	AuthUserId string `json:"auth_user_id"`
}

func Login(serviceName string, authType string) error {
	if authType == "BEARER_TOKEN" {
		bearerLogin(serviceName)
	} else {
		return errors.New("AUTH_TYPE_INVALID")
	}

	return nil
}

func FetchAuthParameters(serviceName string, authType string, inParams *[]rest.InParams) error {
	if authType == "BEARER_TOKEN" {
		bearerFetchAuthParameters(serviceName, inParams)
	} else {
		return errors.New("AUTH_TYPE_INVALID")
	}

	return nil
}
