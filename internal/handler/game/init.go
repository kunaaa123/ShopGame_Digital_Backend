package game

import (
	"github.com/kunaaa123/SHOPGAME/internal/configs"
	"github.com/kunaaa123/SHOPGAME/internal/service/game"
)

type gameHandlerImpl struct {
	service game.GameServiceServer
}

func NewGameHandler(config *configs.Config) gameHandlerImpl {
	service := game.NewGameServiceServer(config)
	return gameHandlerImpl{
		service: service,
	}
}
