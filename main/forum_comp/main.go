package main

import (
	"fmt"
	"github.com/itimofeev/hustledb/forum/prereg"
	"github.com/itimofeev/hustledb/util"
)

func main() {
	r := prereg.ParsePreregCompetition(284, "http://hustle-sa.ru/forum/index.php?showtopic=3753")

	filler := prereg.NewPreregInserter(util.GetDb())
	filler.Insert(r)
	util.PrintJson(r)

	fmt.Println("!!!", "----------") //TODO remove
	f()
}

func f() {
	listLinks := prereg.ParseAllPreregLinks()
	util.PrintJson(listLinks)

	var ids []int
	for _, listLink := range listLinks {
		ids = append(ids, prereg.ParsePreregId(listLink))
	}
	util.PrintJson(ids)

	fCompUrls := make(map[int]string)
	for _, preregId := range ids {
		fCompUrls[preregId] = prereg.GetForumCompetitionId(preregId)
	}

	util.PrintJson(fCompUrls)
}
