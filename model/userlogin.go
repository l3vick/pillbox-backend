package model

type UserLogin struct {
	Mail    	string  `json:"mail,omitempty"`
	Password 	string 	`json:"password,omitempty"`
}