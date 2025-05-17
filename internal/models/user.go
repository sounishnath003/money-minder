package models

type User struct {
	ID       int    `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
