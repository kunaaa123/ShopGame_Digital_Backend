package game

import (
	"gorm.io/gorm"
)

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) (*GameRepository, error) {
	return &GameRepository{db: db}, nil
}
