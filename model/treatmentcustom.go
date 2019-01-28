package model

type TreatmentCustom struct {
	ID              int			`json:"id"`
	IDUser   		int			`json:"id_user"`
	IDMed           int			`json:"id_med"`
	Alarm 			byte		`json:"alarm"`
	Time      		string		`json:"time"`
	StartTreatment  string		`json:"start_treatment"`
	EndTreatment    string		`json:"end_treatment"`
	Period    		int			`json:"period"`
}