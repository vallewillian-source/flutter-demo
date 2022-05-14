package io

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/vallewillian-source/go-sofa-data-studio/models"
)

func Request_params(in_params *[]models.In_params) {

	for i, s := range *in_params {

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
