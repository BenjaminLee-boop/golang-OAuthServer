package models

type User struct {
	Id       string `json:"id"`
	Firtname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
