package model

type UserResponse struct {
	Users []*User `json:"users,omitempty"`
	Page  Page    `json:"page,omitempty"`
}

type UserResponseByPharmacy struct {
	Users []*UserByPharmacy `json:"users,omitempty"`
	Page  Page              `json:"page,omitempty"`
}
