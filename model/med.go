package model

type Med struct {
	ID int		`json:"id,omitempty"`
	Name 		string 	`json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	PharmacyID  int 	`json:"id_pharmacy,omitempty"`
}
