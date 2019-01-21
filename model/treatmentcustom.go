package model

type TreatmentCustom struct {
	ID              int			`json:"id"`
	IDUser   		int			`json:"id_user"`
	IDMed           int			`json:"id_med"`
	Alarm 			byte		`json:"alarm"`
	Time      		string		`json:"time"`
	EndTreatment    string		`json:"end_treatment"`
}