package forum

import "net/http"
import (
	"github.com/gin-gonic/gin"
	"github.com/itimofeev/hustledb/components/forum/comp"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

func NewForumHandlers(db *runner.DB) *ForumHandlers {
	return &ForumHandlers{service: comp.NewFCompService(db)}
}

type ForumHandlers struct {
	service *comp.FCompService
}

func (h *ForumHandlers) ListCompetitions(c *gin.Context) {
	list := h.service.ListCompetitions()
	c.JSON(http.StatusOK, list)
}

func (h *ForumHandlers) ParseCompetitions(c *gin.Context) {
	h.service.ParseCompetitions()
}
