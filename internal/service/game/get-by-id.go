package game

import "github.com/kunaaa123/SHOPGAME/internal/models"

func (g *gameServiceServerImpl) GetByID(id uint) (*models.Game, error) {
	return g.repo.GetByID(id)
}
