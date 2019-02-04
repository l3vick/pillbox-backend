package model

type Treatment struct {
	ID              int			`json:"id,omitempty"`
	IDMed           int			`json:"id_med,omitempty"`
	IDUser   		int			`json:"id_user,omitempty"`
	Morning     	bool		`json:"morning,omitempty"`
	Afternoon      	bool		`json:"afternoon,omitempty"`
	Evening 		bool		`json:"evening,omitempty"`
	StartTreatment  string		`json:"start_treatment,omitempty"`
	EndTreatment    string		`json:"end_treatment,omitempty"`
}

type TreatmentDB struct {
	ID              int			`json:"id,omitempty"`
	IDMed           int			`json:"id_med,omitempty"`
	IDUser   		int			`json:"id_user,omitempty"`
	Morning     	int			`json:"morning"`
	Afternoon      	int			`json:"afternoon"`
	Evening 		int			`json:"evening"`
	StartTreatment  string		`json:"start_treatment,omitempty"`
	EndTreatment    string		`json:"end_treatment,omitempty"`
}
