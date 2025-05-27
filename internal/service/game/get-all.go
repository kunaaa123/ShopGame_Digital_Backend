package game

import "github.com/kunaaa123/SHOPGAME/internal/models"

func (g *gameServiceServerImpl) GetAll() (*[]models.Game, error) {
	return g.repo.GetAll()
}
