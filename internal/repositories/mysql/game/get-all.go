package game

import (
	"github.com/kunaaa123/SHOPGAME/internal/models"
)

func (r *GameRepository) GetAll() (*[]models.Game, error) {
	var data []models.GameDB
	result := r.db.Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	var res []models.Game
	for _, item := range data {
		res = append(res, *item.ToDomain())
	}
	return &res, nil
}
