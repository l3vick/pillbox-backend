package model

type Timing struct {
	Id_User         int		`json:"id_user,omitempty"`
	Morning         byte	`json:"morning,omitempty"`
	Afternoon       byte	`json:"afternoon,omitempty"`
	Evening     	byte	`json:"evening,omitempty"`
	Morning_Time    string 	`json:"morning_time,omitempty"`
	Afternoon_Time	string 	`json:"afternoon_time,omitempty"`
	Evening_Time	string 	`json:"evening_time,omitempty"`
}
