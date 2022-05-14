package models

type In_params struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	Is_required bool   `json:"is_required"`
	Result      string
}
