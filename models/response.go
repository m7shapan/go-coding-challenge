package models

type Response struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload,omitempty"`

	Error interface{} `json:"error,omitempty"`

	CurrentPage int `json:"current_page,omitempty"`
	LastPage    int `json:"last_page,omitempty"`
	PerPage     int `json:"per_page,omitempty"`
	Total       int `json:"total,omitempty"`
}
