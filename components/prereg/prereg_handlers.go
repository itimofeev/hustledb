package prereg

import (
	"github.com/gin-gonic/gin"

	"github.com/itimofeev/hustledb/components/util"
	"net/http"
)

func ListPreregs(c *gin.Context) {
	c.JSON(http.StatusOK, preregService.ListPreregs())
}

func GetPreregById(c *gin.Context) {
	id := c.Param("fCompId")
	prereg := preregService.GetPreregById(util.Atoi(id))
	if prereg == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, prereg)
	}
}

func ParsePreregInfo(c *gin.Context) {
	preregService.ParsePreregInfo()
	c.JSON(http.StatusOK, "OK")
}
