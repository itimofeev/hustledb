package main

import (
	"github.com/itimofeev/hustledb/forum/prereg"
	"github.com/itimofeev/hustledb/util"
)

func main() {
	r := prereg.ParsePreregCompetition(284, "http://hustle-sa.ru/forum/index.php?showtopic=3753")

	filler := prereg.NewPreregFiller(util.GetDb())
	filler.Fill(r)
	util.PrintJson(r)
}

func f() {
	listLinks := prereg.ParseAllPreregLinks()
	util.PrintJson(listLinks)

	var ids []int
	for _, listLink := range listLinks {
		ids = append(ids, prereg.ParsePreregId(listLink))
	}
	util.PrintJson(ids)

	var fCompUrls []string
	for _, preregId := range ids {
		fCompUrl := prereg.GetForumCompetitionId(preregId)
		fCompUrls = append(fCompUrls, fCompUrl)
	}

	util.PrintJson(fCompUrls)
}
