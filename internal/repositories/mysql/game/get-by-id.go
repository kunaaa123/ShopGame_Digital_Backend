package game

import (
	"github.com/kunaaa123/SHOPGAME/internal/models"
)

func (r *GameRepository) GetByID(id uint) (*models.Game, error) {
	m := &models.GameDB{}
	result := r.db.First(&m, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return m.ToDomain(), result.Error
}
