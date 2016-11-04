package main

import (
	"github.com/itimofeev/hustledb/server"
	"github.com/itimofeev/hustledb/util"
	"log"
	"net/http"
)

func main() {
	db := util.GetDb()

	//res := parser.Parse("config.App().JsonFilesPath")
	//parser.InsertData(db, res)

	log.Fatal(http.ListenAndServe(":8080", server.InitRouter(db)))

}
