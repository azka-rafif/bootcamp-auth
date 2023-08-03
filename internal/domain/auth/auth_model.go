package auth

type AuthPayload struct {
	UserName string `json:"userName"`
	Passowrd string `json:"password"`
}

type JwtResponseFormat struct {
	AccessToken string `json:"access_token"`
}
