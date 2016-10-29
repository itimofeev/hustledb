package main

import (
	"encoding/json"
	"fmt"
	"github.com/itimofeev/hustlesa/forum"
	"github.com/itimofeev/hustlesa/util"
	"io/ioutil"
)

const forumDir = "/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustlesa/tools/forum/"

func main() {
	//downloadUrlToFile("http://hustle-sa.ru/forum/index.php?showtopic=2969", forumDir + "2969.html")

	data, err := ioutil.ReadFile(forumDir + "2969.html")
	util.CheckErr(err, "")

	res := forum.GetMainContentFromForumHtml(data)
	mainTitle := forum.GetMainTitleFromForumHtml(data)

	results := forum.ParseForum([]byte(res), mainTitle)

	db := util.GetDb()

	filler := forum.NewForumDbFiller(forum.NewDao(db))
	filler.FillDbInfo(results)

	inserter := forum.NewDbInserter(db)
	inserter.Insert(results)

	jsonData, err := json.Marshal(results)
	util.CheckErr(err, "")

	fmt.Println("!!!", string(jsonData)) //TODO remove
}

func downloadUrlToFile(url, path string) {
	data := forum.GetUrlContent(url)

	ioutil.WriteFile(path, data, 0644)
}

//err := ioutil.WriteFile("/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustlesa/tools/forum/3761.html", data, 0644)
