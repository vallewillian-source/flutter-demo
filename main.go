package main

import (
	"os"

	"github.com/vallewillian-source/go-sofa-data-studio/login"
)

func main() {

	args := os.Args[1:]

	if args[0] == "login_bearer" {
		login.Bearer_login("login_bearer_1.json")
	}

}
