package model

type UserResponse struct {
	Users []*User `json:"users"`
	Page  Page    `json:"page"`
}

type UserResponseByPharmacy struct {
	Users []*UserByPharmacy `json:"users"`
	Page  Page              `json:"page"`
}
