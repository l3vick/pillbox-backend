package model

type PharmacyResponse struct {
	Pharmacy []*Pharmacy `json:"pharmacies,omitempty"`
	Page Page `json:"page,omitempty"`
}