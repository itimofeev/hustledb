package comp

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/itimofeev/hustledb/components/util"
	"regexp"
	"strings"
	"time"
)

const baseUrl = "http://hustle-sa.ru/forum/index.php?showforum=6&prune_day=100&sort_by=Z-A&sort_key=title&st=%d"
const countOnPage = 15
const forumDir = "/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustledb/tools/forum-comp/"

func ParseCompetitionListFromForum() []FCompetition {
	var content []FCompetition
	for page := 0; page < 1000; page += countOnPage {
		url := fmt.Sprintf(baseUrl, page)
		data := util.GetUrlContent(url)
		//data := util.DownloadUrlToFileIfNotExists(url, fmt.Sprintf("%s%d.html", forumDir, page))
		comps := getCompetitionListFromPage(data)
		if len(comps) == 0 {
			return content
		}
		content = append(content, comps...)
	}

	return content
}

func getCompetitionListFromPage(body []byte) []FCompetition {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	util.CheckErr(err)
	r, err := regexp.Compile("^\\(\\d{4}-\\d{2}-\\d{2}(,\\d+)*\\) .*$")
	util.CheckErr(err)

	var comps []FCompetition

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

			comps = append(comps, FCompetition{
				Url:            fixLink(link),
				Title:          compTitle,
				RawText:        s.Text(),
				RawTextChanged: s.Text(),
				HasChange:      true,
				Date:           util.ParseForumDate(compDateStr),
				Desc:           spanElem.Text(),
				ApprovedASH:    hasApprovedStatus(spanElem.Text()),
				City:           tryFindCity(spanElem.Text()),
				DownloadDate:   time.Now(),
			})
		})

	return comps
}

func fixLink(link string) string {
	sIndex := strings.Index(link, "?s=") + 1
	ampIndex := strings.Index(link, "\u0026")

	return link[:sIndex] + link[ampIndex+1:]
}

func parseTitleAndDate(titleAndDate string) (string, string) {
	closeIndex := strings.Index(titleAndDate, ")")

	return titleAndDate[closeIndex+2:], titleAndDate[1:closeIndex]
}

func hasApprovedStatus(desc string) bool {
	lowerDesc := strings.ToLower(desc)
	return strings.Contains(lowerDesc, "утверждено")
}

func tryFindCity(desc string) *string {
	cityIndex := strings.Index(desc, "г.")
	if cityIndex < 0 {
		return nil
	}

	afterCityStart := desc[cityIndex+3:]
	endCityIndex := strings.IndexAny(afterCityStart, ",.")

	if endCityIndex < 0 {
		endCityIndex = len(afterCityStart)
	}

	city := strings.TrimSpace(afterCityStart[:endCityIndex])
	city = strings.Replace(city, "УТВЕРЖДЕНО", "", 1)
	city = strings.Replace(city, "АСХ", "", 1)
	city = strings.Replace(city, "РК", "", 1)
	if len(city) == 0 {
		return nil
	}
	return &city
}
