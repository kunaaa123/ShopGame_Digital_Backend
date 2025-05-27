package game

import (
	"github.com/go-playground/validator/v10"
	"github.com/kunaaa123/SHOPGAME/internal/models"
)

func (g *gameServiceServerImpl) Update(id uint, src models.Game) (*models.Game, error) {
	validate := validator.New()
	err := validate.Struct(src)
	if err != nil {
		return nil, err
	}

	src.ID = id
	err = g.repo.Update(id, src)
	return &src, err
}
