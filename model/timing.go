package model

type Timing struct {
	Id_User         int		`json:"id_user"`
	Morning         bool	`json:"morning"`
	Afternoon       bool	`json:"afternoon"`
	Evening     	bool	`json:"evening"`
	Morning_Time    string 	`json:"morning_time"`
	Afternoon_Time	string 	`json:"afternoon_time"`
	Evening_Time	string 	`json:"evening_time"`
}
