package main

import (
	"github.com/itimofeev/hustlesa/parser"
	"github.com/itimofeev/hustlesa/util"
)

func main() {
	util.InitEnvironment()
	config := util.ReadConfig()
	db := util.InitDb(config)

	res := parser.Parse(config.App().JsonFilesPath)
	parser.InsertData(db, res)
}
