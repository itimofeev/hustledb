package prereg

import (
	"github.com/gin-gonic/gin"
	"github.com/itimofeev/hustledb/forum/prereg"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"

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
