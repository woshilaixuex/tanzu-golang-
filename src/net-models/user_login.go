package net_models

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Vercode  string `json:"vercode"`
}
