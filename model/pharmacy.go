package model

type Pharmacy struct {
	ID              int		`json:"id"`
	Cif            	string	`json:"cif"`
	Street   		string	`json:"street"`
	NumberPhone     int		`json:"number_phone"`
	Schedule       	string	`json:"schedule"`
	Name 			string	`json:"name"`
	Guard     		int		`json:"guard"`
	Password		string	`json:"password"`
	Account			string	`json:"account"`
}