package repository

import (
	"github.com/kunaaa123/SHOPGAME/internal/domain"
	"gorm.io/gorm"
)

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) domain.GameRepository {
	return &GameRepository{db: db}
}

func (r *GameRepository) GetAll() ([]domain.Game, error) {
	var games []domain.Game
	result := r.db.Find(&games)
	return games, result.Error
}

func (r *GameRepository) GetByID(id uint) (domain.Game, error) {
	var game domain.Game
	result := r.db.First(&game, id)
	return game, result.Error
}

func (r *GameRepository) Create(game *domain.Game) error {
	return r.db.Create(game).Error
}

func (r *GameRepository) Update(game *domain.Game) error {
	return r.db.Save(game).Error
}

func (r *GameRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Game{}, id).Error
}
