package model

type User struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Surname     string  `json:"surname,omitempty"`
	Familyname  string  `json:"familyname,omitempty"`
	Password    *string `json:"password,omitempty"`
	Age         int     `json:"age,omitempty"`
	Address     string  `json:"address,omitempty"`
	PhoneNumber int     `json:"phone_number,omitempty"`
	Gender      string  `json:"gender,omitempty"`
	Mail        string  `json:"mail,omitempty"`
	IDPharmacy  int     `json:"id_pharmacy,omitempty"`
	Zip         string  `json:"zip,omitempty"`
	Province    string  `json:"province,omitempty"`
	City        string  `json:"city,omitempty"`
}

type UserResponse struct {
	User     *User           `json:"user,omitempty"`
	Response RequestResponse `json:"response,omitempty"`
}

type UsersResponse struct {
	Users    []*User         `json:"users,omitempty"`
	Page     Page            `json:"page,omitempty"`
	Response RequestResponse `json:"response,omitempty"`
}
