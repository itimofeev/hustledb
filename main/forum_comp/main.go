package main

import (
	"fmt"
	"github.com/itimofeev/hustledb/forum"
	"github.com/itimofeev/hustledb/util"
)

const baseUrl = "http://hustle-sa.ru/forum/index.php?showforum=6&prune_day=100&sort_by=Z-A&sort_key=title&st=%d"
const countOnPage = 15
const forumDir = "/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustledb/tools/forum-comp/"

func main() {
	util.PrintJson(parseCompetitions())
}

func parseCompetitions() []forum.LinkAndTitle {
	var content []forum.LinkAndTitle
	for page := 0; page < 1000; page += countOnPage {
		url := fmt.Sprintf(baseUrl, page)
		data := util.DownloadUrlToFileIfNotExists(url, fmt.Sprintf("%s%d.html", forumDir, page))
		comps := forum.GetCompetitionsFromPage(data)
		if len(comps) == 0 {
			return content
		}
		content = append(content, comps...)
	}

	return content
}
