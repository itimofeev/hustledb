package prereg

import (
	"github.com/gin-gonic/gin"
	"github.com/itimofeev/hustledb/forum/prereg"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"

	"github.com/itimofeev/hustledb/util"
	"net/http"
)

func NewPreregHandlers(db *runner.DB) *PreregHandlers {
	return &PreregHandlers{service: prereg.NewPreregService(db)}
}

type PreregHandlers struct {
	service *prereg.PreregService
}

func (ph *PreregHandlers) ListPreregs(c *gin.Context) {
	c.JSON(http.StatusOK, ph.service.ListPreregs())
}

func (ph *PreregHandlers) GetPreregById(c *gin.Context) {
	id := c.Param("fCompId")
	c.JSON(http.StatusOK, ph.service.GetPreregById(util.Atoi(id)))
}

func (ph *PreregHandlers) ParsePreregInfo(c *gin.Context) {
	ph.service.ParsePreregInfo()
	c.JSON(http.StatusOK, "OK")
}
