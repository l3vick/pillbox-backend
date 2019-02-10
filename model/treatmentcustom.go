package model

type TreatmentCustom struct {
	ID             int    `json:"id,omitempty"`
	IDUser         int    `json:"id_user,omitempty"`
	IDMed          int    `json:"id_med,omitempty"`
	Alarm          string `json:"alarm,omitempty"`
	Time           string `json:"time,omitempty"`
	StartTreatment string `json:"start_treatment,omitempty"`
	EndTreatment   string `json:"end_treatment,omitempty"`
	Period         int    `json:"period,omitempty"`
}
