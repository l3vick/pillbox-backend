package model

type PharmacyResponse struct {
	Pharmacy []*PharmacySql `json:"pharmacies"`
	Page Page `json:"page"`
}