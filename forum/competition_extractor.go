package forum

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/itimofeev/hustledb/util"
	"regexp"
	"strings"
	"fmt"
)

const baseUrl = "http://hustle-sa.ru/forum/index.php?showforum=6&prune_day=100&sort_by=Z-A&sort_key=title&st=%d"
const countOnPage = 15
const forumDir = "/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustledb/tools/forum-comp/"


type LinkAndTitle struct {
	Link    string
	Title   string
	DateStr string
	Desc    string
}

func ParseCompetitionsFromForum() []LinkAndTitle {
	var content []LinkAndTitle
	for page := 0; page < 1000; page += countOnPage {
		url := fmt.Sprintf(baseUrl, page)
		data := util.GetUrlContent(url)
		//data := util.DownloadUrlToFileIfNotExists(url, fmt.Sprintf("%s%d.html", forumDir, page))
		comps := getCompetitionsFromPage(data)
		if len(comps) == 0 {
			return content
		}
		content = append(content, comps...)
	}

	return content
}

func getCompetitionsFromPage(body []byte) []LinkAndTitle {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	util.CheckErr(err)
	r, err := regexp.Compile("^\\(\\d{4}-\\d{2}-\\d{2}(,\\d+)*\\) .*$")
	util.CheckErr(err)

	var lat []LinkAndTitle

	doc.
		Find(".tableborder table tr").
		FilterFunction(func(i int, s *goquery.Selection) bool {
			if s.Find("td").Size() == 0 {
				return false
			}
			if s.Find("th").Size() > 0 {
				return false
			}
			if s.Children().Size() != 7 {
				return false
			}
			if strings.Contains(s.Text(), "ОРГАНИЗАТОРАМ!") {
				return false
			}

			return true
		}).
		Find("td").
		FilterFunction(func(i int, s *goquery.Selection) bool {
			link := s.Find("a")
			if link.Size() != 1 {
				return false
			}

			return r.MatchString(link.Text())
		}).
		Each(func(i int, s *goquery.Selection) {
			aElem := s.Find("a")
			compTitle, compDateStr := parseTitleAndDate(aElem.Text())
			link := aElem.AttrOr("href", "")
			spanElem := s.Find("span")

			lat = append(lat, LinkAndTitle{
				Link:    fixLink(link),
				Title:   compTitle,
				DateStr: compDateStr,
				Desc:    spanElem.Text(),
			})

		})

	return lat
}

func fixLink(link string) string{
	sIndex := strings.Index(link, "?s=") + 1
	ampIndex := strings.Index(link, "\u0026")

	return link[:sIndex] + link[ampIndex+1:]
}

func parseTitleAndDate(titleAndDate string) (string, string) {
	closeIndex := strings.Index(titleAndDate, ")")

	return titleAndDate[closeIndex + 2:], titleAndDate[1:closeIndex]
}
