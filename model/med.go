package model

import (
	"github.com/l3vick/go-pharmacy/nullsql"
)

type Med struct {
	ID int		`json:"id"`
	Name 		string 	`json:"name"`
	Description string  `json:"description"`
	PharmacyID  int 	`json:"id_pharmacy"`
}

type MedSql struct {
	ID * nullsql.JsonNullInt64 `json:"id,omitempty"`
	Name * nullsql.JsonNullString `json:"name,omitempty"`
	Description * nullsql.JsonNullString `json:"description,omitempty"`
	PharmacyID * nullsql.JsonNullInt64 `json:"id_pharmacy,omitempty"`
}
