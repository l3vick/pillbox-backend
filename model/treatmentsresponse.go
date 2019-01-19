package model


type TreatmentsResponse struct {
	Morning 				[]*Morning
	Afternoon				[]*Afternoon
	Evening					[]*Evening
	TreatmentsCustom   		[]*TreatmentCustomResponse	
	Timing           		Timing
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