package auth

import (
	"errors"

	"github.com/vallewillian-source/sofa-lab/internal/rest"
)

func Login(serviceName string, authType string) error {
	if authType == "BEARER_TOKEN" {
		bearerLogin(serviceName)
	} else {
		return errors.New("LOGIN_AUTH_TYPE_INVALID")
	}

	return nil
}

func FetchAuthParameters(serviceName string, authType string, inParams *[]rest.InParams) error {
	if authType == "BEARER_TOKEN" {
		bearerFetchAuthParameters(serviceName, inParams)
	} else {
		return errors.New("FETCH_AUTH_TYPE_INVALID")
	}

	return nil
}
