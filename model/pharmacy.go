package model

type Pharmacy struct {
	ID              int		`json:"id"`
	Cif            	string	`json:"cif"`
	Street   		int		`json:"street"`
	NumberPhone     int		`json:"number_phone"`
	Schedule       	int		`json:"schedule"`
	Name 			string	`json:"name"`
	Guard     		string	`json:"guard"`
	Password		string	`json:"password"`
}