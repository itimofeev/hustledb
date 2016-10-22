package forum

import (
	"bufio"
	"bytes"
	"github.com/itimofeev/hustlesa/util"
	"regexp"
)

var (
	judgeRegexp       *regexp.Regexp
	participantRegexp *regexp.Regexp
)

func initRegexps() {
	judgeRegexp = compileRegexp("\\d \\([A-G]\\) - [\\W ]+")                        //1 (A) - Милованов Александр
	participantRegexp = compileRegexp(".* ((Участников)||(Участвовало пар)): \\d+") //DnD Beginner (ПАРТНЕРЫ). Участников: 49 |  DnD Beginner (ДЕВУШКИ). Участников: 85 | E класс. Участвовало пар: 23
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
		state = state.ProcessLine(results, curLine)
	}

	return results
}

type BeginState struct {
}
type JudgeTeamState struct {
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
			return &BeginState{}
		}
	default:
		return s
	}
}

type FAState interface {
	ProcessLine(fr *ForumResults, line string) FAState
}
