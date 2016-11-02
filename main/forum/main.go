package main

import (
	"github.com/itimofeev/hustlesa/forum"
	"github.com/itimofeev/hustlesa/util"
	"io/ioutil"
)

const forumDir = "/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustlesa/tools/forum/"

const compUrl = "http://hustle-sa.ru/forum/index.php?showtopic="

func main() {
	//parseAndInsert("2969") //(2014-09-06) Открытие сезона (г.Москва), ДК Буревестник, м.Сокольники
	//parseAndInsert("2929") //(2014-09-20,21) Hustle & Discofox Festival Cup, В рамках H&D RUSSIAN OPEN FESTIVAL
	//parseAndInsert("3007") //(2014-09-27) Восходящие звезды, г. Санкт-Петербург
	//parseAndInsert("3000") //(2014-09-28) Отрытый турнир для Е+D классов, г. Красноярск
	//parseAndInsert("2988") //(2014-10-04) Огни большого города, г.Москва
	//parseAndInsert("2981") //(2014-10-11) "Перезагрузка" (ТК"Движение"), г. Новосибирск, УТВЕРЖДЕН
	//parseAndInsert("3029") //(2014-10-18) Танцевальный Weekend 2014, г. Санкт-Петербург, Утверждено
	//parseAndInsert("3039") //(2014-10-25) Турнир "Красный октябрь", г. Москва, Утверждено РК АСХ
	//parseAndInsert("3058") //(2014-10-25) Шаг Вперед, г.Красноярск УТВЕРЖДЕНО
	parseAndInsert("3048") //(2014-10-26) Первенство Поволжья, УТВЕРЖДЕНО, г. Саратов
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
