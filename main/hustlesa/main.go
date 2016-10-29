package main

import (
	"github.com/itimofeev/hustlesa/server"
	"github.com/itimofeev/hustlesa/util"
	"log"
	"net/http"
)

func main() {
	db := util.GetDb()

	//res := parser.Parse("config.App().JsonFilesPath")
	//parser.InsertData(db, res)

	log.Fatal(http.ListenAndServe(":8080", server.InitRouter(db)))

}
