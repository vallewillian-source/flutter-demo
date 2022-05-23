package models

type Scheema struct {
	Name       string          `json:"name"`
	PrimaryKey string          `json:"primary_key"`
	Fields     []ScheemaFields `json:"fields"`
}

type ScheemaFields struct {
	Name    string `json:"name"`
	Scheema string `json:"scheema"`
	Address string `json:"address"`
}
