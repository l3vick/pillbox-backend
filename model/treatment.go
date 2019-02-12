package model

type Treatment struct {
	ID             int    `json:"id,omitempty"`
	IDMed          int    `json:"id_med,omitempty"`
	IDUser         int    `json:"id_user,omitempty"`
	Morning        string `json:"morning,omitempty"`
	Afternoon      string `json:"afternoon,omitempty"`
	Evening        string `json:"evening,omitempty"`
	StartTreatment string `json:"start_treatment,omitempty"`
	EndTreatment   string `json:"end_treatment,omitempty"`
}

type TreatmentResponse struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Morning        string `json:"morning,omitempty"`
	Afternoon      string `json:"afternoon,omitempty"`
	Evening        string `json:"evening,omitempty"`
	StartTreatment string `json:"start_treatment,omitempty"`
	EndTreatment   string `json:"end_treatment,omitempty"`
}

type TreatmentsResponse struct {
	TreatmentsMorning   []*TreatmentResponse       `json:"treatments_morning,omitempty"`
	TreatmentsAfternoon []*TreatmentResponse       `json:"treatments_afternoon,omitempty"`
	TreatmentsEvening   []*TreatmentResponse       `json:"treatments_evening,omitempty"`
	TreatmentsCustom    []*TreatmentCustomResponse `json:"treatments_custom,omitempty"`
	Timing              Timing                     `json:"timing,omitempty"`
	Response            []RequestResponse          `json:"response,omitempty"`
}
