package model

type UserResponse struct {
	Users []*User 	`json:"users"`
	Page  Page 		`json:"page"`
}

type UserResponseNotNull struct {
	Users []*UserNotNull 	`json:"users"`
	Page  Page 				`json:"page"`
}
