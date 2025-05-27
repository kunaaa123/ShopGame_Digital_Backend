package game

import "github.com/kunaaa123/SHOPGAME/internal/models"

func (r *GameRepository) Create(src models.Game) error {
	result := r.db.Create(src.ToDB())
	if result.Error != nil {
		return result.Error
	}
	return nil
}
