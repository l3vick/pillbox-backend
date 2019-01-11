package model

type UserLogin struct {
	Phone    int    `json:"number_phone"`
	Password string `json:"password"`
}