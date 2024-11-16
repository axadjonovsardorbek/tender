package models

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"` // "client", "admin"
}

type Login struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type Token struct {
    Token string `json:"token"`
}