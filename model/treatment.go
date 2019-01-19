package model

type Treatment struct {
	ID              int			`json:"id"`
	IDMed           int			`json:"id_med"`
	IDUser   		int			`json:"id_user"`
	Morning     	byte		`json:"morning"`
	Afternoon      	byte		`json:"afternoon"`
	Evening 		byte		`json:"evening"`
	EndTreatment    string		`json:"end_treatment"`
}
