package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (g *gameHandlerImpl) GetGameByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	res, err := g.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, res)
}
