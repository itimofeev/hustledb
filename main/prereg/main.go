package main

import (
	"github.com/itimofeev/hustledb/forum/prereg"
	"github.com/itimofeev/hustledb/util"
)

func main() {
	r := prereg.ParsePreregCompetition(284, "http://hustle-sa.ru/forum/index.php?showtopic=3753")

	inserter := prereg.NewPreregInserter(util.GetDb())
	inserter.Insert(r)
	util.PrintJson(r)

}

func f() {
	listLinks := prereg.ParseAllPreregLinks()
	util.PrintJson(listLinks)

	var ids []int
	for _, listLink := range listLinks {
		ids = append(ids, prereg.ParsePreregId(listLink))
	}

	fCompUrls := make(map[int]string)
	for _, preregId := range ids {
		fCompUrls[preregId] = prereg.GetForumCompetitionId(preregId)
	}
}
