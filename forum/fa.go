package forum

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/itimofeev/hustlesa/util"
	"regexp"
)

var (
	judgeRegexp       *regexp.Regexp
	participantRegexp *regexp.Regexp
	stageRegexp       *regexp.Regexp
	stageFinalRegexp  *regexp.Regexp
	placeRegexp       *regexp.Regexp
)

func initRegexps() {
	judgeRegexp = compileRegexp("\\d \\([A-G]\\) - [\\W ]+")                        //1 (A) - Милованов Александр
	participantRegexp = compileRegexp(".* ((Участников)||(Участвовало.пар)):.\\d+") //DnD Beginner (ПАРТНЕРЫ). Участников: 49 |  DnD Beginner (ДЕВУШКИ). Участников: 85 | E класс. Участвовало пар: 23
	stageRegexp = compileRegexp("1/\\d+ финала")                                    //1/2 финала | ФИНАЛ | 1/16 финала
	stageFinalRegexp = compileRegexp("ФИНАЛ")                                       //1/2 финала | ФИНАЛ | 1/16 финала
	placeRegexp = compileRegexp("\\d+(-\\d+)? место-№\\d+-.*")                      //7-11 место-№510-Потехин Алексей Викторович(6117,AlphaDance,D)-Степнова Наталья Андреевна(6398,AlphaDance,D) | 6 место-№553-Фадеев Алексей Сергеевич(8599,Движение,D,Bg)
}

func compileRegexp(rx string) *regexp.Regexp {
	compiled, err := regexp.Compile(rx)
	util.CheckErr(err, "Unable to compile: "+rx)
	return compiled
}

func ParseForum(data []byte) *ForumResults {
	initRegexps()

	scanner := bufio.NewScanner(bytes.NewReader(data))

	results := &ForumResults{}
	var state FAState = &BeginState{}

	for scanner.Scan() {
		curLine := scanner.Text()
		fmt.Printf("State: %T, line: '%s'\n", state, curLine) //TODO remove
		state = state.ProcessLine(results, curLine)
	}

	return results
}

type BeginState struct {
}
type JudgeTeamState struct {
}
type PlacesState struct {
}

func (s *BeginState) ProcessLine(fr *ForumResults, line string) FAState {
	switch line {
	case "Результаты турнира:":
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
		return &JudgeTeamState{}
	}

	return s
}

type FAState interface {
	ProcessLine(fr *ForumResults, line string) FAState
}
