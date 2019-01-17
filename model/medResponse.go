package model

type MedResponse struct {
	Meds []*MedSql `json:"meds"`
	Page Page `json:"page"`
}
