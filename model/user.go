package model

import "github.com/l3vick/go-pharmacy/util"

type User struct {
	ID              int		`json:"id"`
	Name            string	`json:"name"`
	SurName        	string	`json:"surname"`
	FamilyName     	string	`json:"familyname"`
	Password       	string 	`json:"password"`
	Age				int 	`json:"age"`
	Address			string 	`json:"address"`
	Phone			int 	`json:"phone_number"`
	Gender			string 	`json:"gender"`
	Mail			string 	`json:"mail"`
	IDPharmacy     	int    	`json:"id_pharmacy"`
	Zip				string 	`json:"zip"`
	Province		string 	`json:"province"`
	City     		string  `json:"city"`
}

type UserNotNull struct {
	ID              int						`json:"id"`
	Name            string					`json:"name"`
	SurName         string					`json:"surname"`
	FamilyName     	util.JsonNullString		`json:"familyname,omitempty"`
	Password *      util.JsonNullString 	`json:"password,omitempty"`
	Age				util.JsonNullInt64 		`json:"age,omitempty"`
	Address			util.JsonNullString 	`json:"address,omitempty"`
	Phone			util.JsonNullInt64 		`json:"phone_number,omitempty"`
	Gender			util.JsonNullString 	`json:"gender,omitempty"`
	Mail			util.JsonNullString 	`json:"mail,omitempty"`
	IDPharmacy     	int    					`json:"id_pharmacy"`
	Zip				util.JsonNullString 	`json:"zip,omitempty"`
	Province		util.JsonNullString 	`json:"province,omitempty"`
	City     		util.JsonNullString  	`json:"city,omitempty"`
}
