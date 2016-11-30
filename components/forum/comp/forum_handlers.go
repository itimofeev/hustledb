package comp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCompetitions(c *gin.Context) {
	list := GetFCompService().ListCompetitions()
	c.JSON(http.StatusOK, list)
}

func ParseCompetitions(c *gin.Context) {
	GetFCompService().ParseCompetitions()
}
