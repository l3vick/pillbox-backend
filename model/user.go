package model

//User ...
type User struct {
	ID              int		`json:"id"`
	Name            string	`json:"name"`
	MedBreakfast   	string	`json:"med_breakfast"`
	MedLaunch       string	`json:"med_launch"`
	MedDinner       string	`json:"med_dinner"`
	AlarmBreakfast 	string	`json:"alarm_breakfast"`
	AlarmLaunch     string	`json:"alarm_launch"`
	AlarmDinner     string	`json:"alarm_dinner"`
	Password       	string 	`json:"password"`
	IDPharmacy     	int    	`json:"id_pharmacy"`
}
