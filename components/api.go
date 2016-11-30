package components

import (
	"github.com/gin-gonic/gin"
	fServer "github.com/itimofeev/hustledb/components/forum/results"
	"github.com/itimofeev/hustledb/components/hsaxls"
	"github.com/itimofeev/hustledb/components/prereg"
	"github.com/itimofeev/hustledb/components/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"net/http"
)

func InitRouter(conn *runner.DB, session *mgo.Session) *gin.Engine {
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

	fHandlers := fServer.NewForumHandlers(conn)
	forumApi := api.Group("/forum")
	forumApi.GET("/competitions", fHandlers.ListCompetitions)
	forumApi.POST("/competitions", fHandlers.ParseCompetitions)

	preregHandlers := prereg.NewPreregHandlers(conn, session)
	preregApi := api.Group("/prereg")
	preregApi.GET("/", preregHandlers.ListPreregs)
	preregApi.GET("/:fCompId", preregHandlers.GetPreregById)
	preregApi.POST("/", preregHandlers.ParsePreregInfo)

	util.DB = conn

	return r
}
