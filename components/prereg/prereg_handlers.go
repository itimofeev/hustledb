package prereg

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"

	"github.com/itimofeev/hustledb/components/util"
	"gopkg.in/mgo.v2"
	"net/http"
)

func NewPreregHandlers(db *runner.DB, session *mgo.Session) *PreregHandlers {
	return &PreregHandlers{service: NewPreregService(db, session)}
}

type PreregHandlers struct {
	service *PreregService
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
