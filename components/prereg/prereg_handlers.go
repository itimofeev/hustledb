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
	c.JSON(http.StatusOK, preregService.GetPreregById(util.Atoi(id)))
}

func ParsePreregInfo(c *gin.Context) {
	preregService.ParsePreregInfo()
	c.JSON(http.StatusOK, "OK")
}
