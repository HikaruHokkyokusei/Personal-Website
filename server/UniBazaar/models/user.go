package models

type User struct {
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	IsVerified  bool   `json:"is_verified"`
	Affiliation string `json:"affiliation"`
}
