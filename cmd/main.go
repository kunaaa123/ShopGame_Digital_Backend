package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kunaaa123/SHOPGAME/internal/adapter/handler"
	"github.com/kunaaa123/SHOPGAME/internal/adapter/repository"
	"github.com/kunaaa123/SHOPGAME/internal/db"
	"github.com/kunaaa123/SHOPGAME/internal/service"
)

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	dbConfig := db.NewDBConfig()
	database := db.GetDB(dbConfig)

	gameRepo := repository.NewGameRepository(database)
	gameService := service.NewgameService(gameRepo)
	gameHandler := handler.NewGameHandler(gameService)

	r := gin.Default()
	gameHandler.RegisterRoutes(r)

	log.Println("Server starting on http://localhost:8080")
	r.Run(":8080")
}
