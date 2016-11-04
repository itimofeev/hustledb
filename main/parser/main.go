package main

import (
	"github.com/itimofeev/hustledb/parser"
	"github.com/itimofeev/hustledb/util"
)

func main() {
	util.InitEnvironment()
	config := util.ReadConfig()
	db := util.InitDb(config)

	res := parser.Parse(config.App().JsonFilesPath)
	parser.InsertData(db, res)
}
