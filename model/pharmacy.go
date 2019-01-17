package model

import "github.com/l3vick/go-pharmacy/nullsql"

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

type PharmacySql struct {
	ID              nullsql.JsonNullInt64	`json:"id"`
	Cif            	*nullsql.JsonNullString	`json:"cif,omitempty"`
	Address   		*nullsql.JsonNullString	`json:"address,omitempty"`
	NumberPhone     *nullsql.JsonNullInt64	`json:"number_phone,omitempty"`
	Schedule       	*nullsql.JsonNullString	`json:"schedule,omitempty"`
	Name 			*nullsql.JsonNullString	`json:"name,omitempty"`
	Guard     		*nullsql.JsonNullInt64	`json:"guard,omitempty"`
	Password		*nullsql.JsonNullString	`json:"password,omitempty"`
	Mail			*nullsql.JsonNullString	`json:"mail,omitempty"`
}