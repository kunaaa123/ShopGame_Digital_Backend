package models

type Game struct {
	ID          uint    `json:"id,omitempty"`
	Name        string  `json:"name,omitempty" validate:"required"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty" validate:"required"`
	Image       string  `json:"image,omitempty" validate:"required"`
}
