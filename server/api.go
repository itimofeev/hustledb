package server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"net/http"
	"os"
)

var db *runner.DB

func InitRouter(conn *runner.DB) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api/v1").Subrouter() /*.Headers("Content-Type", "application/json")*/

	compRouter := apiRouter.PathPrefix("/competitions").Subrouter()
	dancerRouter := apiRouter.PathPrefix("/dancers").Subrouter()

	compRouter.Methods("GET").HandlerFunc(ListCompetitions)
	dancerRouter.Methods("GET").HandlerFunc(ListDancers)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	recovery := RecoveryHandler(PrintRecoveryStack(true))(loggedRouter)

	db = conn

	return recovery
}

func ListCompetitions(w http.ResponseWriter, r *http.Request) {
	var params PageParams
	parsePageParams(w, r, &params)

	t := RepoListCompetitions(params)

	WriteJSONStatus(w, t, http.StatusOK)
}

func ListDancers(w http.ResponseWriter, r *http.Request) {
	var params PageParams
	parseParamsGet(w, r, &params)

	t := RepoListDancers(params)

	WriteJSONStatus(w, t, http.StatusOK)
}
