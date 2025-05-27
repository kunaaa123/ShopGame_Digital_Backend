package game

import (
	"github.com/go-playground/validator/v10"
	"github.com/kunaaa123/SHOPGAME/internal/models"
)

func (g *gameServiceServerImpl) Create(src models.Game) (*models.Game, error) {
	validate := validator.New()
	err := validate.Struct(src)
	if err != nil {
		return nil, err
	}

	err = g.repo.Create(src)
	return &src, err
}
