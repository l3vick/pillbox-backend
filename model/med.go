package model

type Med struct {
	ID int		`json:"id"`
	Name 		string 	`json:"name"`
	Description string  `json:"description"`
	PharmacyID  int 	`json:"id_pharmacy"`
}
