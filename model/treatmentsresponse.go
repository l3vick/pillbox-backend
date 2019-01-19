package model


type TreatmentsResponse struct {
	Morning 				[]*Morning
	Afternoon				[]*Afternoon
	Evening					[]*Evening
	TreatmentsCustom   		[]*TreatmentCustomResponse	
	Timing           		TimingResponse
}

type Morning struct {
	ID			int
	Name		string
}

type Afternoon struct {
	ID			int
	Name		string
}

type Evening struct {
	ID			int
	Name		string
}

type TreatmentCustomResponse struct {
	ID			int
	Name		string
	Time		string
	Alarm		byte
}

type TimingResponse struct {
	Morning         byte	`json:"morning"`
	Afternoon       byte	`json:"afternoon"`
	Evening     	byte	`json:"evening"`
	Morning_Time    string 	`json:"morning_time"`
	Afternoon_Time	string 	`json:"afternoon_time"`
	Evening_Time	string 	`json:"evening_time"`
}