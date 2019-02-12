package model

type Med struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IDPharmacy  int    `json:"id_pharmacy,omitempty"`
}
type MedsResponse struct {
	Meds     []*Med          `json:"meds,omitempty"`
	Page     Page            `json:"page,omitempty"`
	Response RequestResponse `json:"response,omitempty"`
}

type MedResponse struct {
	Med      *Med            `json:"meds,omitempty"`
	Response RequestResponse `json:"response,omitempty"`
}
