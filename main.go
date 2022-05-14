package main

import (
	"fmt"
	"os"

	"github.com/vallewillian-source/go-sofa-data-studio/login"
)

func main() {

	args := os.Args[1:]
	fmt.Println(args)

	if args[0] == "login_bearer" {
		login.Bearer_login("login_bearer_1.json")
	}

}
