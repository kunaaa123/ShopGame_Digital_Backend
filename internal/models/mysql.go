package models

import (
	"gorm.io/gorm"
)

type GameDB struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Image       string
}
