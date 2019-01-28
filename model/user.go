package model

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	SurName    string `json:"surname"`
	FamilyName string `json:"familyname"`
	Password   string `json:"password"`
	Age        int    `json:"age"`
	Address    string `json:"address"`
	Phone      int    `json:"phone_number"`
	Gender     string `json:"gender"`
	Mail       string `json:"mail"`
	IDPharmacy int    `json:"id_pharmacy"`
	Zip        string `json:"zip"`
	Province   string `json:"province"`
	City       string `json:"city"`
}

type UserByPharmacy struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	SurName    string `json:"surname"`
	FamilyName string `json:"familyname"`
	Age        int    `json:"age"`
	Address    string `json:"address"`
	Phone      int    `json:"phone_number"`
	Gender     string `json:"gender"`
	Mail       string `json:"mail"`
	Zip        string `json:"zip"`
	Province   string `json:"province"`
	City       string `json:"city"`
}
