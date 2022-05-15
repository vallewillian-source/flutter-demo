package models

type Endpoint struct {
	Name         string       `json:"name"`
	Url          string       `json:"url"`
	Method       string       `json:"method"`
	Auth_service string       `json:"auth_service"`
	Auth_type    string       `json:"auth_type"`
	In_params    []In_params  `json:"in_params"`
	Out_params   []Out_params `json:"out_params"`
}

type Out_params struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Scheema string `json:"scheema"`
}
