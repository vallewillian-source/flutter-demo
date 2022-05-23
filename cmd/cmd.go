package cmd

import (
	"os"

	"github.com/vallewillian-source/go-sofa-data-studio/internal/auth"
)

func Execute() {
	args := os.Args[1:]

	if args[0] == "login_bearer" {
		auth.BearerLogin(args[1])
	} else if args[0] == "run_endpoint" {
		Run(args[1])
	}
}
