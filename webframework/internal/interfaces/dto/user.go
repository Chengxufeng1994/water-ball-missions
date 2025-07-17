package dto

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (u *User) Key() string {
	return "User"
}
