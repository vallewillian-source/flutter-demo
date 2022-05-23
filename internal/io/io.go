package io

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/vallewillian-source/sofa-lab/internal/rest"
)

func FetchParams(in_params *[]rest.InParams) {

	for i, s := range *in_params {
		// TODO implement post and querystring
		if s.Type == "body" {
			prompt := promptui.Prompt{
				Label: s.Name,
			}
			result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
			} else {
				(*in_params)[i].Result = result
			}
		}

	}

}
