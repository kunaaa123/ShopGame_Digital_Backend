package game

import (
	"github.com/kunaaa123/SHOPGAME/internal/models"
)

func (r *GameRepository) Update(id uint, src models.Game) error {
	m := &models.GameDB{}
	result := r.db.Model(m).Where(id).Updates(src.ToDB())
	if result.Error != nil {
		return result.Error
	}
	return nil
}
