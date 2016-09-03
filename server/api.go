package server

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/itimofeev/hustlesa/model"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var db *runner.DB

func InitRouter(conn *runner.DB) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api/v1"). /*.Headers("Content-Type", "application/json")*/
							Subrouter()

	compRouter := apiRouter.PathPrefix("/competitions").Subrouter()

	compRouter.Methods("GET").HandlerFunc(ListCompetitions)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	recovery := RecoveryHandler(PrintRecoveryStack(true))(loggedRouter)

	db = conn

	return recovery
}

func ListCompetitions(w http.ResponseWriter, r *http.Request) {
	var params PageParams
	parsePageParams(w, r, &params)

	t, err := RepoListCompetitions(params)
	if err != nil {
		panic(err)
	}

	WriteJSONStatus(w, t, http.StatusCreated)
}

func RepoListCompetitions(params PageParams) ([]model.RawCompetition, error) {
	var competitions []model.RawCompetition
	err := db.SQL(`
		SELECT
			*
		FROM
			competition
		LIMIT $1
		OFFSET $2
	`,
		params.Limit,
		params.Offset,
	).QueryStructs(&competitions)
	return competitions, err
}

func parsePageParams(w http.ResponseWriter, r *http.Request, params *PageParams) {
	parseParams(w, r, params)

	if params.Limit > MaxLimit {
		params.Limit = MaxLimit
	}

	if params.Limit == 0 {
		params.Limit = DefaultLimit
	}
}

func parseParams(w http.ResponseWriter, r *http.Request, params interface{}) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if len(body) == 0 {
		return
	}

	if err := json.Unmarshal(body, &params); err != nil {
		panic(err)
	}
}
