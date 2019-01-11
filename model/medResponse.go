package model

type MedResponse struct {
	Meds []*Med `json:"meds"`
	Page Page `json:"page"`
}
