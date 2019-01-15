package model

import (
	"github.com/l3vick/go-pharmacy/util"
)

type Med struct {
	ID   int	`json:"id"`
	Name string `json:"name"`
	Description  string    `json:"description"`
	PharmacyID  util.JsonNullInt64 `json:"pharmacy_id,omitempty"`
}

type MedInt struct {
	ID   int	`json:"id"`
	Name string `json:"name"`
	Description  string    `json:"description"`
	PharmacyID  int `json:"pharmacy_id,omitempty"`
}