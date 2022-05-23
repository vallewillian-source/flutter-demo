package models

type BearerLoginService struct {
	ServiceName   string              `json:"service_name"`
	AuthType      string              `json:"auth_type"`
	LoginEndpoint BearerLoginEndpoint `json:"login_endpoint"`
}

type BearerLoginEndpoint struct {
	Url       string               `json:"url"`
	Method    string               `json:"method"`
	InParams  []InParams           `json:"in_params"`
	OutParams BearerLoginOutParams `json:"out_params"`
}

type BearerLoginOutParams struct {
	AuthToken  string `json:"auth_token"`
	AuthUserId string `json:"auth_user_id"`
}

type BearerLoginFile struct {
	AuthToken  string `json:"auth_token"`
	AuthUserId string `json:"auth_user_id"`
}
