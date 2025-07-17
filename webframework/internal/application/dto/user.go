package dto

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

type Token struct {
	Token string `json:"token"`
}
