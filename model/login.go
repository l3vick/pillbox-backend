package model

type LoginUser struct {
	Mail     string `json:"mail,omitempty"`
	Password string `json:"password,omitempty"`
}

type CheckMail struct {
	ID    int    `json:"id,omitempty"`
	Mail  string `json:"mail,omitempty"`
	State bool   `json:"state,omitempty"`
}

type LoginResponse struct {
	Pharmacy *Pharmacy       `json:"pharmacy,omitempty"`
	Response RequestResponse `json:"response,omitempty"`
}

type CheckMailResponse struct {
	CheckMail *CheckMail      `json:"pharmacy,omitempty"`
	Response  RequestResponse `json:"response,omitempty"`
}
