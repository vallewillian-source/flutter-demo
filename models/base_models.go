package models

type InParams struct {
	Address    string `json:"address"`
	Name       string `json:"name"`
	IsRequired bool   `json:"is_required"`
	Result     string
	Type       string `json:"type"`
	Auth       string `json:"auth"`
}
