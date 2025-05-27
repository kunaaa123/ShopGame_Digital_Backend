package game

import "github.com/gin-gonic/gin"

type GameHandler interface {
	GetAllGames(c *gin.Context)
	GetGameByID(c *gin.Context)
	CreateGame(c *gin.Context)
	UpdateGame(c *gin.Context)
	DeleteGame(c *gin.Context)
}
