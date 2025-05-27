package game

import (
	"log"

	"github.com/kunaaa123/SHOPGAME/internal/configs"
	"github.com/kunaaa123/SHOPGAME/internal/repositories"
	"github.com/kunaaa123/SHOPGAME/internal/repositories/mysql"
	"github.com/kunaaa123/SHOPGAME/internal/repositories/mysql/game"
)

type gameServiceServerImpl struct {
	repo repositories.GameRepository
}

func NewGameServiceServer(configs *configs.Config) GameServiceServer {
	if configs == nil {
		panic("configs is nil")
	}

	db := mysql.GetDB(configs)
	repo, err := game.NewGameRepository(db)
	if err != nil {
		log.Panic("cannot initialize game repositorie ", err)
	}

	return &gameServiceServerImpl{
		repo: repo,
	}
}
