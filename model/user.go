package model

type User struct {
	ID              int		`json:"id"`
	Name            string	`json:"name"`
	Surname         string	`json:"surname"`
	SurnameLast     string	`json:"surname_last"`
	MedBreakfast   	string	`json:"med_breakfast"`
	MedLaunch       string	`json:"med_launch"`
	MedDinner       string	`json:"med_dinner"`
	AlarmBreakfast 	string	`json:"alarm_breakfast"`
	AlarmLaunch     string	`json:"alarm_launch"`
	AlarmDinner     string	`json:"alarm_dinner"`
	Password       	string 	`json:"password"`
	Age				int 	`json:"age"`
	Address			string 	`json:"address"`
	Phone			int 	`json:"phone"`
	Gender			string 	`json:"gender"`
	Mail			string 	`json:"mail"`
	IDPharmacy     	int    	`json:"id_pharmacy"`
}
