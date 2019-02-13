package model

type Timing struct {
	IDUser        int    `json:"id_user,omitempty"`
	Morning       string `json:"morning,omitempty"`
	Afternoon     string `json:"afternoon,omitempty"`
	Evening       string `json:"evening,omitempty"`
	MorningTime   string `json:"morning_time,omitempty"`
	AfternoonTime string `json:"afternoon_time,omitempty"`
	EveningTime   string `json:"evening_time,omitempty"`
}

type TimingResponse struct {
	Timing   Timing          `json:"timing,omitempty"`
	Response RequestResponse `json:"response,omitempty"`
}
