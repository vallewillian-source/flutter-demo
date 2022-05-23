package models

type Endpoint struct {
	Name        string      `json:"name"`
	Url         string      `json:"url"`
	Method      string      `json:"method"`
	AuthService string      `json:"auth_service"`
	AuthType    string      `json:"auth_type"`
	InParams    []InParams  `json:"in_params"`
	OutParams   []OutParams `json:"out_params"`
}

type OutParams struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Scheema string `json:"scheema"`
}
