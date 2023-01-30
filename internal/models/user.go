package models

type UserJSON struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegistrationResponse struct {
	Result bool `json:"result"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
