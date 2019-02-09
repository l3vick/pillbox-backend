package model

type User struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	SurName    string `json:"surname,omitempty"`
	FamilyName string `json:"familyname,omitempty"`
	Password   string `json:"password,omitempty"`
	Age        int    `json:"age,omitempty"`
	Address    string `json:"address,omitempty"`
	Phone      int    `json:"phone_number,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Mail       string `json:"mail,omitempty"`
	IDPharmacy int    `json:"id_pharmacy,omitempty"`
	Zip        string `json:"zip,omitempty"`
	Province   string `json:"province,omitempty"`
	City       string `json:"city,omitempty"`
}

type UserByPharmacy struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	SurName    string `json:"surname,omitempty"`
	FamilyName string `json:"familyname,omitempty"`
	Age        int    `json:"age,omitempty"`
	Address    string `json:"address,omitempty"`
	Phone      int    `json:"phone_number,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Mail       string `json:"mail,omitempty"`
	Zip        string `json:"zip,omitempty"`
	Province   string `json:"province,omitempty"`
	City       string `json:"city,omitempty"`
}
