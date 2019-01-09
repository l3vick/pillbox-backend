package model

//User ...
type User struct {
	ID              int		`json:"id"`
	Name            string	`json:"name"`
	MedBreakfast   int		`json:"med_breakfast"`
	MedLaunch       int		`json:"med_launch"`
	MedDinner       int		`json:"med_dinner"`
	AlarmBreakfast string	`json:"alarm_breakfast"`
	AlarmLaunch     string	`json:"alarm_launch"`
	AlarmDinner     string	`json:"alarm_dinner"`
	Password       	string 	`json:"password"`
	IDPharmacy     	int    	`json:"id_pharmacy"`
}
