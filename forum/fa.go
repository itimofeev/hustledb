package forum

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/itimofeev/hustlesa/util"
	"regexp"
	"strings"
	"sync"
	"time"
)

var once = sync.Once{}

var (
	judgeRegexp          *regexp.Regexp
	participantRegexp    *regexp.Regexp
	stageRegexp          *regexp.Regexp
	stageFinalRegexp     *regexp.Regexp
	techStageRegexp      *regexp.Regexp
	techFinalStageRegexp *regexp.Regexp
	placeRegexp          *regexp.Regexp
	techResult           *regexp.Regexp
	techFinalResult      *regexp.Regexp
	tableBorder          *regexp.Regexp
	resultsRegexp        *regexp.Regexp
	mainTitleRegexp      *regexp.Regexp
)

func initRegexps() {
	judgeRegexp = compileRegexp("\\d \\(.\\) - [\\W ]+") //1 (A) - Милованов Александр
	resultsRegexp = compileRegexp("\\s*Результаты турнира[,\\d\\W ]*:")
	participantRegexp = compileRegexp(".* ((Участников)||(Участвовало.пар)):.\\d+") //DnD Beginner (ПАРТНЕРЫ). Участников: 49 |  DnD Beginner (ДЕВУШКИ). Участников: 85 | E класс. Участвовало пар: 23
	stageRegexp = compileRegexp("1/\\d+ финала")                                    //1/2 финала | ФИНАЛ | 1/16 финала
	stageFinalRegexp = compileRegexp("ФИНАЛ")                                       //1/2 финала | ФИНАЛ | 1/16 финала
	placeRegexp = compileRegexp("\\d+(-\\d+)? место-№\\d+-.*")                      //7-11 место-№510-Потехин Алексей Викторович(6117,AlphaDance,D)-Степнова Наталья Андреевна(6398,AlphaDance,D) | 6 место-№553-Фадеев Алексей Сергеевич(8599,Движение,D,Bg)

	techStageRegexp = compileRegexp(".*: 1/\\d+ финала")                        // D класс: 1/8 финала | DnD Beginner (ДЕВУШКИ): 1/16 финала
	techFinalStageRegexp = compileRegexp(".*ФИНАЛ")                             // DnD Beginner (ДЕВУШКИ): ФИНАЛ
	techResult = compileRegexp(".*\\d+.*\\|.*\\|.*((==> выход в)||(место:)).*") //   502   | AB|CDE ==> выход в 1/8 финала || 579   | AD|   место: 23-30
	techFinalResult = compileRegexp("^.*\\d+.*│(.*\\d.*)+│(.*│)+.*\\d+$")       //  687   | 1 5 4 2 6         │     5  │ 4 4 3 2 5         │     4  │                   │        │    9 │    4

	tableBorder = compileRegexp("(-+\\+)+-+") //--------+-------------------+--------+-------------------+--------+-------------------+--------+------+---------

	mainTitleRegexp = compileRegexp("\\(\\d{4}-\\d{2}-\\d{2}\\) [\\W\\s.\\-\\d()]+,.*") //(2016-09-17) Открытие сезона 2016-2017г., г. Москва. УТВЕРЖДЕНО РК АСХ || (2014-09-06) Открытие сезона (г.Москва), ДК Буревестник, м.Сокольники
}

func compileRegexp(rx string) *regexp.Regexp {
	compiled, err := regexp.Compile(rx)
	util.CheckErr(err, "Unable to compile: "+rx)
	return compiled
}

func ParseForum(data []byte, mainTitle string) *ForumResults {
	once.Do(initRegexps)

	scanner := bufio.NewScanner(bytes.NewReader(data))

	results := &ForumResults{}
	var state FAState = &BeginState{}

	for scanner.Scan() {
		curLine := scanner.Text()
		fmt.Printf("State: %T, line: '%s'\n", state, curLine) //TODO remove
		state = state.ProcessLine(results, curLine)
	}

	parseMainTitle(results, mainTitle)

	return results
}

//(2016-09-17) Открытие сезона 2016-2017г., г. Москва. УТВЕРЖДЕНО РК АСХ
//(2014-09-06) Открытие сезона (г.Москва), ДК Буревестник, м.Сокольники
func parseMainTitle(results *ForumResults, mainTitle string) {
	once.Do(initRegexps)

	mainTitle = strings.Replace(mainTitle, ", УТВЕРЖДЕНО РК АСХ", ". УТВЕРЖДЕНО РК АСХ", 1)
	mainTitle = strings.TrimSpace(mainTitle)
	util.CheckMatchesRegexp(mainTitleRegexp.String(), mainTitle)

	closeIndex := strings.Index(mainTitle, ")")
	dateStr := mainTitle[1:closeIndex]

	layout := "2006-01-02"
	date, err := time.Parse(layout, dateStr)
	util.CheckErr(err, "Time parse")
	results.Date = date

	commaIndex := strings.Index(mainTitle, ", ")
	results.Title = strings.TrimSpace(mainTitle[closeIndex+1 : commaIndex])

	results.Remaining = mainTitle[commaIndex+2:]
}

type BeginState struct {
}
type JudgeTeamState struct {
}
type PlacesState struct {
}
type TechnicalState struct {
}
type TechnicalPrepFinalState struct {
}
type TechnicalFinalState struct {
}

func (s *BeginState) ProcessLine(fr *ForumResults, line string) FAState {
	once.Do(initRegexps)

	switch {
	case resultsRegexp.MatchString(line):
		return &JudgeTeamState{}
	default:
		return s
	}
}

func (s *JudgeTeamState) ProcessLine(fr *ForumResults, line string) FAState {
	switch {
	case judgeRegexp.MatchString(line): //1 (A) - Милованов Александр
		{
			fr.addJudge(line)
			return &JudgeTeamState{}
		}
	case participantRegexp.MatchString(line): // E класс. Участвовало пар: 23 | DnD Beginner (ПАРТНЕРЫ). Участников: 49 |  DnD Beginner (ДЕВУШКИ). Участников: 85
		{
			fr.addNominationName(line)
			return &PlacesState{}
		}
	default:
		return s
	}

	return nil
}

func (s *PlacesState) ProcessLine(fr *ForumResults, line string) FAState {
	switch {
	case stageRegexp.MatchString(line) || stageFinalRegexp.MatchString(line): //1/2 финала | ФИНАЛ | 1/16 финала
		{
			fr.addStage(line)
			return s
		}
	case placeRegexp.MatchString(line): //7-11 место-№510-Потехин Алексей Викторович(6117,AlphaDance,D)-Степнова Наталья Андреевна(6398,AlphaDance,D) | 6 место-№553-Фадеев Алексей Сергеевич(8599,Движение,D,Bg)
		{
			fr.addPlace(line)
			return s
		}
	case line == "Технические результаты:":
		return &TechnicalState{}
	}

	return s
}

func (s *TechnicalState) ProcessLine(fr *ForumResults, line string) FAState {
	switch {
	case techStageRegexp.MatchString(line): // D класс: 1/8 финала | DnD Beginner (ДЕВУШКИ): 1/16 финала
		{
			fr.addTechStage(line)
			return s
		}
	case techResult.MatchString(line): //   502   | AB|CDE ==> выход в 1/8 финала || 579   | AD|   место: 23-30
		{
			fr.addTechResult(line)
			return s
		}
	case techFinalStageRegexp.MatchString(line): // DnD Beginner (ДЕВУШКИ): ФИНАЛ
		{
			fr.addFinalTechStage(line)
			return &TechnicalPrepFinalState{}
		}
	}

	return s
}

func (s *TechnicalPrepFinalState) ProcessLine(fr *ForumResults, line string) FAState {
	switch {
	case techFinalResult.MatchString(line): //  687   | 1 5 4 2 6         │     5  │ 4 4 3 2 5         │     4  │                   │        │    9 │    4
		{
			fr.addFinalTechResult(line)
			return &TechnicalFinalState{}
		}
	}

	return s
}

func (s *TechnicalFinalState) ProcessLine(fr *ForumResults, line string) FAState {
	switch {
	case tableBorder.MatchString(line):
		{
			return &JudgeTeamState{}
		}
	case techFinalResult.MatchString(line): //  687   | 1 5 4 2 6         │     5  │ 4 4 3 2 5         │     4  │                   │        │    9 │    4
		{
			fr.addFinalTechResult(line)
			return s
		}
	}

	return s
}

type FAState interface {
	ProcessLine(fr *ForumResults, line string) FAState
}
