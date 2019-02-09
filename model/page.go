package model

type Page struct {
	First int `json:"first,omitempty"`
	Previous int `json:"previous,omitempty"`
	Next int `json:"next,omitempty"`
	Last  int `json:"last,omitempty"`
	Count int `json:"count,omitempty"`
}
