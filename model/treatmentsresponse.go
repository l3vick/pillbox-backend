package model

type TreatmentsResponse struct {
	Morning          []*Morning                 `json:"morning,omitempty"`
	Afternoon        []*Afternoon               `json:"afternoon,omitempty"`
	Evening          []*Evening                 `json:"evening,omitempty"`
	TreatmentsCustom []*TreatmentCustomResponse `json:"treatments_custom,omitempty"`
	Timing           TimingResponse             `json:"timing,omitempty"`
	Response         []RequestResponse          `json:"response,omitempty"`
}

type Morning struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	StartTreatment string `json:"start_treatment,omitempty"`
	EndTreatment   string `json:"end_treatment,omitempty"`
}

type Afternoon struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	StartTreatment string `json:"start_treatment,omitempty"`
	EndTreatment   string `json:"end_treatment,omitempty"`
}

type Evening struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
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

type TimingResponse struct {
	Morning        bool   `json:"morning,omitempty"`
	Afternoon      bool   `json:"afternoon,omitempty"`
	Evening        bool   `json:"evening,omitempty"`
	Morning_Time   string `json:"morning_time,omitempty"`
	Afternoon_Time string `json:"afternoon_time,omitempty"`
	Evening_Time   string `json:"evening_time,omitempty"`
}
