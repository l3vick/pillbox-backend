package model

type TreatmentsResponse struct {
	TreatmentsMorning   []*TreatmentResponse       `json:"treatments_morning,omitempty"`
	TreatmentsAfternoon []*TreatmentResponse       `json:"treatments_afternoon,omitempty"`
	TreatmentsEvening   []*TreatmentResponse       `json:"treatments_evening,omitempty"`
	TreatmentsCustom    []*TreatmentCustomResponse `json:"treatments_custom,omitempty"`
	Timing              Timing                     `json:"timing,omitempty"`
	Response            []RequestResponse          `json:"response,omitempty"`
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

type TreatmentCustomResponse struct {
	ID             int    `json:"id,omitempty"`
	IDMed          int    `json:"id_med,omitempty"`
	Name           string `json:"name,omitempty"`
	Time           string `json:"time,omitempty"`
	Alarm          string `json:"alarm,omitempty"`
	StartTreatment string `json:"start_treatment,omitempty"`
	EndTreatment   string `json:"end_treatment,omitempty"`
	Period         int    `json:"period,omitempty"`
}
