package repositories

import "github.com/kunaaa123/SHOPGAME/internal/models"

type GameRepository interface {
	GetAll() (*[]models.Game, error)
	GetByID(id uint) (*models.Game, error)
	Create(src models.Game) error
	Update(id uint, src models.Game) error
	Delete(id uint) error
}
