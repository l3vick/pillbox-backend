package model

import "github.com/l3vick/go-pharmacy/util"

type Pharmacy struct {
	ID              int		`json:"id"`
	Cif            	string	`json:"cif"`
	Address   		string	`json:"address"`
	NumberPhone     int		`json:"number_phone"`
	Schedule       	string	`json:"schedule"`
	Name 			string	`json:"name"`
	Guard     		int		`json:"guard"`
	Password		string	`json:"password"`
	Mail			string	`json:"mail"`
}

type PharmacyNotNull struct {
	ID              int					`json:"id"`
	Cif            	util.JsonNullString	`json:"cif,omitempty"`
	Address   		util.JsonNullString	`json:"address,omitempty"`
	NumberPhone     util.JsonNullInt64	`json:"number_phone,omitempty"`
	Schedule       	util.JsonNullString	`json:"schedule,omitempty"`
	Name 			util.JsonNullString	`json:"name,omitempty"`
	Guard     		util.JsonNullInt64	`json:"guard,omitempty"`
	Password *		util.JsonNullString	`json:"password,omitempty"`
	Mail			util.JsonNullString	`json:"mail,omitempty"`
}