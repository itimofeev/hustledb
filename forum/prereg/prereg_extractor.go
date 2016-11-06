package prereg

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/itimofeev/hustledb/util"
	"strings"
)

type PreregComp struct {
	ID            int64              `json:"id" db:"id"`
	PreregId      int                `json:"prereg_id" db:"prereg_id"`
	CompetitionId string             `json:"competition_id" db:"competition_id"`
	Nominations   []PreregNomination `json:"nominations"`
}

type PreregNomination struct {
	ID           int64  `json:"id" db:"id"`
	PreregCompId int64  `json:"prereg_comp_id" db:"prereg_comp_id"`
	Title        string `json:"title" db:"title"`
	Records      string `json:"records"`
}

type PreregRecord struct {
	ID          int64 `json:"id" db:"id"`
	PreregNomId int64 `json:"prereg_nom_id" db:"prereg_nom_id"`
	Index       int   `json:"index" db:"idnex"`
}

type PreregDancer struct {
	ID      int64        `json:"id" db:"id"`
	CodeASH *string      `json:"code_ash" db:"code_ash"`
	Class   string       `json:"class" db:"class"`
	Title   string       `json:"title" db:"title"`
	Clubs   []PreregClub `json:"clubs"`
}

type PreregClub struct {
	ID    int64  `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}

const mainUrl = "http://www.liveindance.ru/"
const regUrl = "http://www.liveindance.ru/contest/reg.php?id=%d"

func ParseAllPreregLinks() []string {
	data := util.GetUrlContent(mainUrl)
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(data))
	util.CheckErr(err)

	var listLinks []string

	doc.
		Find(".panel-body a").
		FilterFunction(func(i int, s *goquery.Selection) bool {
			link := s.AttrOr("href", "")
			return strings.Contains(link, "http://www.liveindance.ru/contest/registration/list.php?id=")
		}).
		Each(func(i int, s *goquery.Selection) {
			link := s.AttrOr("href", "")
			listLinks = append(listLinks, link)
		})
	return listLinks
}

func ParsePreregId(listLink string) int {
	preregId := strings.Replace(listLink, "http://www.liveindance.ru/contest/registration/list.php?id=", "", 1)

	return util.Atoi(preregId)
}

func GetForumCompetitionId(preregId int) string {
	regUrlFull := fmt.Sprintf(regUrl, preregId)
	data := util.GetUrlContent(regUrlFull)

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(data))
	util.CheckErr(err)

	forumLinkA := doc.
		Find("table table table tr a").
		FilterFunction(func(i int, s *goquery.Selection) bool {
			return strings.Contains(s.AttrOr("href", ""), "http://hustle-sa.ru/forum/index.php?showtopic=")
		}).First()

	return forumLinkA.AttrOr("href", "")
}

func ParsePreregCompetition(preregId int, fCompUrl string) PreregComp {
	return PreregComp{}
}
