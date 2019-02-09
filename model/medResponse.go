package model

type MedResponse struct {
	Meds []*Med `json:"meds,omitempty"`
	Page Page `json:"page,omitempty"`
}
