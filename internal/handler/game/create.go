package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunaaa123/SHOPGAME/internal/models"
)

func (g *gameHandlerImpl) CreateGame(c *gin.Context) {
	var body models.Game
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := g.service.Create(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
