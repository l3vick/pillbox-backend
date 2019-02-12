package model

type UserResponse struct {
	User     *User           `json:"users,omitempty"`
	Response RequestResponse `json:"response,omitempty"`
}

type UsersResponse struct {
	Users    []*User         `json:"users,omitempty"`
	Page     Page            `json:"page,omitempty"`
	Response RequestResponse `json:"response,omitempty"`
}
