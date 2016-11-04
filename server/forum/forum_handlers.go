package forum

import "net/http"
import (
	"github.com/gin-gonic/gin"
	"github.com/itimofeev/hustledb/forum"
)

func NewForumHandlers(dao forum.FCompDao) *ForumHandlers {
	return &ForumHandlers{dao: dao}
}

type ForumHandlers struct {
	dao forum.FCompDao
}

func (h *ForumHandlers) ListCompetitions(c *gin.Context) {
	list := h.dao.ListCompetitions()
	c.JSON(http.StatusOK, list)
}

func (h *ForumHandlers) ParseCompetitions(c *gin.Context) {
	linkAndTitles := forum.ParseCompetitionsFromForum()
	for _, value := range linkAndTitles {
		h.dao.CreateCompetition(value)
	}
}
