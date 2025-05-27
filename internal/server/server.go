package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kunaaa123/SHOPGAME/internal/configs"
	"github.com/kunaaa123/SHOPGAME/internal/handler/game"
)

func RunServer() error {
	return RunServerWithConfig(configs.New())
}

func RunServerWithConfig(configs *configs.Config) error {
	r := gin.Default()
	handler := game.NewGameHandler(configs)
	app := r.Group("/games")
	{
		app.GET("/", handler.GetAllGames)
		app.GET("/:id", handler.GetGameByID)
		app.POST("/", handler.CreateGame)
		app.PUT("/:id", handler.UpdateGame)
		app.DELETE("/:id", handler.DeleteGame)
	}

	port := configs.Server.RESTPort

	log.Printf("Server starting on http://localhost%s", port)
	r.Run(port)

	return nil
}
