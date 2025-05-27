package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kunaaa123/SHOPGAME/internal/models"
)

func (g *gameHandlerImpl) UpdateGame(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	var src models.Game
	if err := c.ShouldBindJSON(&src); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := g.service.Update(uint(id), src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
