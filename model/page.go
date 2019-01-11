package model

//Page ...
type Page struct {
	First int `json:"first"`
	Previous int `json:"previous"`
	Next int `json:"next"`
	Last  int `json:"last"`
	Count int `json:"count"`
}
