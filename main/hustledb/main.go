package main

import (
	server "github.com/itimofeev/hustledb/components"
	"github.com/itimofeev/hustledb/components/util"

	"log"
	"net/http"
)

func main() {
	util.InitPersistence()
	server.InitCronTasks()

	log.Fatal(http.ListenAndServe(":8080", server.InitRouter()))

}
