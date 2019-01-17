package model

type UserResponse struct {
	Users []*User 	`json:"users"`
	Page  Page 		`json:"page"`
}

type UserResponseSql struct {
	Users []*UserSql 	`json:"users"`
	Page  Page 			`json:"page"`
}
