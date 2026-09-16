package models

type Room struct {
	ID          uint    `json:"id"`
	Class       string  `json:"class"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
