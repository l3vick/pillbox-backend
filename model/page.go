package model

type Page struct {
	First    int `json:"first"`
	Previous int `json:"previous"`
	Next     int `json:"next,omitempty"`
	Last     int `json:"last,omitempty"`
	Count    int `json:"count,omitempty"`
}
