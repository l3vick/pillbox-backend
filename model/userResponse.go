package model

type UserResponse struct {
	Users []*User `json:"users"`
	Page Page `json:"page"`
}
