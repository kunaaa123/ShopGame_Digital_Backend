package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (g *gameHandlerImpl) GetAllGames(c *gin.Context) {
	res, err := g.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
