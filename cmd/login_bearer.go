package cmd

type bearer_login_endpoint struct {
	service_name   string
	auth_type      string
	login_endpoint bearer_login_endpoint_in_params
}

type bearer_login_endpoint_in_params struct {
	address     string
	name        string
	is_required bool
}

func login_bearer() {
	print("OK")
}
