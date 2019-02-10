package model

type Treatment struct {
	ID              int			`json:"id,omitempty"`
	IDMed           int			`json:"id_med,omitempty"`
	IDUser   		int			`json:"id_user,omitempty"`
	Morning     	string		`json:"morning,omitempty"`
	Afternoon      	string		`json:"afternoon,omitempty"`
	Evening 		string		`json:"evening,omitempty"`
	StartTreatment  string		`json:"start_treatment,omitempty"`
	EndTreatment    string		`json:"end_treatment,omitempty"`
}
