package game

import (
	"github.com/kunaaa123/SHOPGAME/internal/models"
)

type GameServiceServer interface {
	GetAll() (*[]models.Game, error)
	GetByID(id uint) (*models.Game, error)
	Create(src models.Game) (*models.Game, error)
	Update(id uint, src models.Game) (*models.Game, error)
	Delete(id uint) error
}
