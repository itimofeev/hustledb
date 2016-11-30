package hsaxls

import (
	"github.com/gin-gonic/gin"
	"github.com/itimofeev/hustledb/components/util"
	"net/http"
)

func ListCompetitions(c *gin.Context) {
	var params util.PageParams
	util.ParseParamsGet(c, &params)
	params.Fix()

	t := RepoListCompetitions(params)

	util.WriteJSONStatus(c, t, http.StatusOK)
}

func ListCompetition(c *gin.Context) {
	compId := util.GetPathInt64Param(c, "id")

	t := RepoGetCompetitionInfo(compId)

	util.WriteJSONStatus(c, t, http.StatusOK)
}

func ListDancers(c *gin.Context) {
	var params util.ListDancerParams
	util.ParseParamsGet(c, &params)
	params.Fix()

	t := RepoListDancers(params)

	util.WriteJSONStatus(c, t, http.StatusOK)
}

func GetDancerInfo(c *gin.Context) {
	dancerId := util.GetPathInt64Param(c, "dancerId")

	t := RepoGetDancerInfo(dancerId)

	util.WriteJSONStatus(c, t, http.StatusOK)
}
