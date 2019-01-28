package model


type TreatmentsResponse struct {
	Morning 				[]*Morning						`json:"morning"`	
	Afternoon				[]*Afternoon					`json:"afternoon"`
	Evening					[]*Evening						`json:"evening"`
	TreatmentsCustom   		[]*TreatmentCustomResponse		`json:"treatments_custom"`
	Timing           		TimingResponse					`json:"timing"`
}

type Morning struct {
	ID				int		`json:"id"`
	Name			string	`json:"name"`
	StartTreatment	string	`json:"start_treatment"`
	EndTreatment	string	`json:"end_treatment"`
}

type Afternoon struct {
	ID				int		`json:"id"`
	Name			string	`json:"name"`
	StartTreatment	string	`json:"start_treatment"`
	EndTreatment	string	`json:"end_treatment"`
}

type Evening struct {
	ID				int		`json:"id"`
	Name			string	`json:"name"`
	StartTreatment	string	`json:"start_treatment"`
	EndTreatment	string	`json:"end_treatment"`
}

type TreatmentCustomResponse struct {
	ID				int		`json:"id"`
	Name			string	`json:"name"`
	Time			string	`json:"time"`
	Alarm			byte	`json:"alarm"`
	StartTreatment 	string	`json:"start_treatment"`
	EndTreatment	string	`json:"end_treatment"`
	Period			int		`json:"period"`
}

type TimingResponse struct {
	Morning         byte	`json:"morning"`
	Afternoon       byte	`json:"afternoon"`
	Evening     	byte	`json:"evening"`
	Morning_Time    string 	`json:"morning_time"`
	Afternoon_Time	string 	`json:"afternoon_time"`
	Evening_Time	string 	`json:"evening_time"`
}