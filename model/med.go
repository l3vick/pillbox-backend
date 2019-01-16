package model

import (
	"github.com/l3vick/go-pharmacy/util"
)

type Med struct {
	ID   		int					`json:"id"`
	Name 		string 				`json:"name"`
	Description string    			`json:"description"`
	PharmacyID  util.JsonNullInt64 	`json:"id_pharmacy,omitempty"`
}

type MedInt struct {
	ID   		int		`json:"id"`
	Name 		string 	`json:"name"`
	Description string  `json:"description"`
	PharmacyID  int 	`json:"id_pharmacy,omitempty"`
}