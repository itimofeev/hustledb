package main

import (
	"github.com/itimofeev/hustledb/components/prereg"
	"github.com/itimofeev/hustledb/components/util"
)

func main() {
	r := prereg.ParsePreregCompetition(284, "http://hustle-sa.ru/forum/index.php?showtopic=3753")

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
