package components

import (
	"github.com/gin-gonic/gin"
	"github.com/itimofeev/hustledb/components/forum/comp"
	"github.com/itimofeev/hustledb/components/hsaxls"
	"github.com/itimofeev/hustledb/components/prereg"
	"github.com/itimofeev/hustledb/components/util"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.LoggerWithWriter(util.GinLog.Writer()), gin.RecoveryWithWriter(util.RecLog.Writer()))

	api := r.Group("/api/v1")

	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to hustledb!")
	})

	api.GET("/competitions", hsaxls.ListCompetitions)
	api.GET("/competitions/:id", hsaxls.ListCompetition)
	api.GET("/dancers", hsaxls.ListDancers)
	api.GET("/dancers/:dancerId", hsaxls.GetDancerInfo)

	forumApi := api.Group("/forum")
	forumApi.GET("/competitions", comp.ListCompetitions)

	preregApi := api.Group("/prereg")
	preregApi.GET("/", prereg.ListPreregs)
	preregApi.GET("/:fCompId", prereg.GetPreregById)

	admin := api.Group("/admin")
	admin.POST("/competitions", comp.ParseCompetitions)
	admin.POST("/prereg", prereg.ParsePreregInfo)

	return r
}
