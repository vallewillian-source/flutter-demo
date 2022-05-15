package models

type Bearer_login_service struct {
	Service_name   string                `json:"service_name"`
	Auth_type      string                `json:"auth_type"`
	Login_endpoint Bearer_login_endpoint `json:"login_endpoint"`
}

type Bearer_login_endpoint struct {
	Url        string                  `json:"url"`
	Method     string                  `json:"method"`
	In_params  []In_params             `json:"in_params"`
	Out_params Bearer_login_out_params `json:"out_params"`
}

type Bearer_login_out_params struct {
	Auth_token   string `json:"auth_token"`
	Auth_user_id string `json:"auth_user_id"`
}

type Bearer_login_file struct {
	Auth_token   string `json:"auth_token"`
	Auth_user_id string `json:"auth_user_id"`
}
