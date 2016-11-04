package main

import (
	"github.com/itimofeev/hustledb/forum"
	"github.com/itimofeev/hustledb/util"
)

func main() {
	util.PrintJson(forum.ParseCompetitionsFromForum())
}
