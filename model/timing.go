package model

type Timing struct {
	Id_User        int    `json:"id_user"`
	Morning        byte   `json:"morning"`
	Afternoon      byte   `json:"afternoon"`
	Evening        byte   `json:"evening"`
	Morning_Time   string `json:"morning_time"`
	Afternoon_Time string `json:"afternoon_time"`
	Evening_Time   string `json:"evening_time"`
}

type TimingDB struct {
	IDUser        int    `json:"id_user"`
	Morning       bool   `json:"morning"`
	Afternoon     bool   `json:"afternoon"`
	Evening       bool   `json:"evening"`
	MorningTime   string `json:"morning_time"`
	AfternoonTime string `json:"afternoon_time"`
	EveningTime   string `json:"evening_time"`
}
