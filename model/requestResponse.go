package model

type RequestResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
