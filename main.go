package main

import (
	"os"

	"github.com/vallewillian-source/go-sofa-data-studio/login"
	"github.com/vallewillian-source/go-sofa-data-studio/run_endpoint"
)

func main() {

	args := os.Args[1:]

	if args[0] == "login_bearer" {
		login.Bearer_login("login_bearer_1.json")
	} else if args[0] == "run_endpoint" {
		run_endpoint.Execute("endpoint_1.json")
	}

}
