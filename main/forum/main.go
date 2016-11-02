package main

import (
	"github.com/itimofeev/hustlesa/forum"
	"github.com/itimofeev/hustlesa/util"
	"io/ioutil"
)

const forumDir = "/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustlesa/tools/forum/"

const compUrl = "http://hustle-sa.ru/forum/index.php?showtopic="

func main() {
	//parseAndInsert("2969")//(2014-09-06) Открытие сезона (г.Москва), ДК Буревестник, м.Сокольники
	//parseAndInsert("2929")//(2014-09-20,21) Hustle & Discofox Festival Cup, В рамках H&D RUSSIAN OPEN FESTIVAL
	parseAndInsert("3007") //(2014-09-27) Восходящие звезды, г. Санкт-Петербург
}

func parseAndInsert(topicId string) {
	filePath := forumDir + topicId + ".html"
	if !util.IsFileExists(filePath) {
		downloadUrlToFile(compUrl+topicId, filePath)
	}

	data, err := ioutil.ReadFile(filePath)
	util.CheckErr(err, "")

	res := forum.GetMainContentFromForumHtml(data)
	mainTitle := forum.GetMainTitleFromForumHtml(data)

	results := forum.ParseForum([]byte(res), mainTitle)

	util.PrintJson(results)

	db := util.GetDb()

	filler := forum.NewForumDbFiller(compUrl+topicId, forum.NewDao(db))
	filler.FillDbInfo(results)

	inserter := forum.NewDbInserter(forum.NewInsertDao(db))
	inserter.Insert(results)
}

func downloadUrlToFile(url, path string) {
	data := forum.GetUrlContent(url)

	ioutil.WriteFile(path, data, 0644)
}

//err := ioutil.WriteFile("/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustlesa/tools/forum/3761.html", data, 0644)
