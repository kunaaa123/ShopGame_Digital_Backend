package game

import "github.com/kunaaa123/SHOPGAME/internal/models"

func (r *GameRepository) Delete(id uint) error {
	m := &models.GameDB{}
	result := r.db.Where(id).Delete(m)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GameRepository) ForceDelete(id uint) error {
	m := &models.GameDB{}
	result := r.db.Unscoped().Where(id).Delete(m)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
