package models

type Scheema struct {
	Name        string           `json:"name"`
	Primary_key string           `json:"primary_key"`
	Fields      []Scheema_fields `json:"fields"`
}

type Scheema_fields struct {
	Name    string `json:"name"`
	Scheema string `json:"scheema"`
	Address string `json:"address"`
}
