package model

type Pharmacy struct {
	ID          int    `json:"id"`
	Cif         string `json:"cif"`
	Address     string `json:"address"`
	NumberPhone int    `json:"number_phone"`
	Schedule    string `json:"schedule"`
	Name        string `json:"name"`
	Guard       int    `json:"guard"`
	Password    string `json:"password"`
	Mail        string `json:"mail"`
}

type PharmacyR struct {
	ID    int    `json:"id"`
	Mail  string `json:"mail"`
	State bool   `json:"state"`
}
