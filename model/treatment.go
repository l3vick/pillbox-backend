package model

type Treatment struct {
	ID              int			`json:"id"`
	IDMed           int			`json:"id_med"`
	IDUser   		int			`json:"id_user"`
	Morning     	bool		`json:"morning"`
	Afternoon      	bool		`json:"afternoon"`
	Evening 		bool		`json:"evening"`
	EndTreatment    string		`json:"end_treatment"`
}