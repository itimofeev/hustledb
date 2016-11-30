package components

import (
	"github.com/gin-gonic/gin"
	fServer "github.com/itimofeev/hustledb/components/forum/results"
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

	fHandlers := fServer.NewForumHandlers(util.DB)
	forumApi := api.Group("/forum")
	forumApi.GET("/competitions", fHandlers.ListCompetitions)

	preregHandlers := prereg.NewPreregHandlers(util.DB, util.MGO)
	preregApi := api.Group("/prereg")
	preregApi.GET("/", preregHandlers.ListPreregs)
	preregApi.GET("/:fCompId", preregHandlers.GetPreregById)

	admin := api.Group("/admin")
	admin.POST("/competitions", fHandlers.ParseCompetitions)
	admin.POST("/prereg", preregHandlers.ParsePreregInfo)

	return r
}
