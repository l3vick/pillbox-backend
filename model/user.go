package model

import "github.com/l3vick/go-pharmacy/nullsql"

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

type UserSql struct {
	ID              nullsql.JsonNullInt64   `json:"id"`
	Name            *nullsql.JsonNullString  `json:"name,omitempty"`
	SurName        	*nullsql.JsonNullString   `json:"surname,omitempty"`
	FamilyName     	*nullsql.JsonNullString  `json:"familyname,omitempty"`
	Password       	*nullsql.JsonNullString 	`json:"password,omitempty"`
	Age				*nullsql.JsonNullInt64 	`json:"age,omitempty"`
	Address			*nullsql.JsonNullString 	`json:"address,omitempty"`
	Phone			*nullsql.JsonNullInt64 	`json:"phone_number,omitempty"`
	Gender			*nullsql.JsonNullString 	`json:"gender,omitempty"`
	Mail			*nullsql.JsonNullString 	`json:"mail,omitempty"`
	IDPharmacy     	nullsql.JsonNullInt64   `json:"id_pharmacy"`
	Zip				*nullsql.JsonNullString 	`json:"zip,omitempty"`
	Province		*nullsql.JsonNullString 	`json:"province,omitempty"`
	City     		*nullsql.JsonNullString  `json:"city,omitempty"`
}
