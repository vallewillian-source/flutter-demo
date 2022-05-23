package cmd

import (
	"os"
)

type apiFile struct {
	Name     string `json:"name"`
	AuthType string `json:"auth_type"`
}

func Execute() {
	args := os.Args[1:]

	if args[0] == "login" {
		Login(args[1])
	} else if args[0] == "run" {
		result, err := Run(args[1], args[2])
		if err != nil {
			panic(err)
		}

		err = exportToJson(args[1], args[2], result)
		if err != nil {
			panic(err)
		}

	}
}
