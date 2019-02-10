package model

type Pharmacy struct {
	ID          int    `json:"id,omitempty"`
	Cif         string `json:"cif,omitempty"`
	Address     string `json:"address,omitempty"`
	NumberPhone int    `json:"phone_number,omitempty"`
	Schedule    string `json:"schedule,omitempty"`
	Name        string `json:"name,omitempty"`
	Guard       int    `json:"guard,omitempty"`
	Password    string `json:"password,omitempty"`
	Mail        string `json:"mail,omitempty"`
}

type PharmacyR struct {
	ID    int    `json:"id,omitempty"`
	Mail  string `json:"mail,omitempty"`
	State bool   `json:"state,omitempty"`
}
