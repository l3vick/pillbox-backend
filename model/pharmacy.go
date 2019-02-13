package model

type Pharmacy struct {
	ID          int     `json:"id,omitempty"`
	Cif         string  `json:"cif,omitempty"`
	Address     string  `json:"address,omitempty"`
	PhoneNumber int     `json:"phone_number,omitempty"`
	Schedule    string  `json:"schedule,omitempty"`
	Name        string  `json:"name,omitempty"`
	Guard       string  `json:"guard,omitempty"`
	Password    *string `json:"password,omitempty"`
	Mail        string  `json:"mail,omitempty"`
}

type PharmacyR struct {
	ID    int    `json:"id,omitempty"`
	Mail  string `json:"mail,omitempty"`
	State bool   `json:"state,omitempty"`
}

type PharmaciesResponse struct {
	Pharmacies []*Pharmacy     `json:"pharmacies,omitempty"`
	Page       Page            `json:"page,omitempty"`
	Response   RequestResponse `json:"response,omitempty"`
}

type PharmacyResponse struct {
	Pharmacy *Pharmacy       `json:"pharmacy,omitempty"`
	Response RequestResponse `json:"response,omitempty"`
}
