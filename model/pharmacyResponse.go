package model

type PharmacyResponse struct {
	Pharmacy []*Pharmacy `json:"pharmacies"`
	Page Page `json:"page"`
}