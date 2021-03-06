package forum

import (
	"database/sql"
	"time"
)

type ForumResults struct {
	JudgesResults []*JudgeTeam
	Date          time.Time
	Title         string
	Remaining     string

	CompetitionId int64
}

type JudgeTeam struct {
	Judges      []*Judge
	Nominations []*Nomination
}

type Judge struct {
	ID       int64  `db:"id"`
	DancerId int64  `db:"dancer_id"`
	Letter   string `db:"letter"`
	Title    string

	PartitionId int64 `db:"partition_id"`
}

type Nomination struct {
	ID              int64  `db:"id"`
	Title           string `db:"title"`
	Count           int
	Stages          []*Stage
	TechStages      []*TechStage
	FinalTechStage  string
	FinalTechResult []*FinalTechResult

	PartitionId   int64         `db:"partition_id"`
	RNominationId sql.NullInt64 `db:"r_nomination_id"`
}

type Dancer struct {
	Id           int64
	Title        string
	ClassClassic string
	ClassJnj     string
	Clubs        []string
}

type Place struct {
	ID        int64 `db:"id"`
	PlaceFrom int   `db:"place_from"`
	PlaceTo   int   `db:"place_to"`
	Number    int   `db:"number"`
	Dancer1   *Dancer
	Dancer2   *Dancer

	StageTitle   string        `db:"stage_title"`
	NominationId int64         `db:"nomination_id"`
	Dancer1Id    int64         `db:"dancer1_id"`
	Dancer2Id    sql.NullInt64 `db:"dancer2_id"`

	Result1Id sql.NullInt64 `db:"result1_id"`
	Result2Id sql.NullInt64 `db:"result2_id"`
}

type Stage struct {
	Title  string
	Places []*Place
}

type TechStage struct {
	Title   string
	Results []*TechResult
}

type TechResult struct {
	Value string
}
type FinalTechResult struct {
	Value string
}

func (fr *ForumResults) addJudge(line string) {
	judge := parseJudge(line)
	needNew := len(fr.JudgesResults) == 0 || len(fr.JudgesResults[len(fr.JudgesResults)-1].Nominations) > 0
	if needNew {
		fr.JudgesResults = append(fr.JudgesResults, &JudgeTeam{
			Judges: []*Judge{judge},
		})
	} else {
		lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]
		lastJudgeResult.Judges = append(lastJudgeResult.Judges, judge)
	}
}

func (fr *ForumResults) addNominationName(line string) {
	lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]

	title, count := parseNominationTitleCount(line)

	lastJudgeResult.Nominations = append(lastJudgeResult.Nominations, &Nomination{
		Title: title,
		Count: count,
	})
}

func (fr *ForumResults) addStage(line string) {
	lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]
	lastNomination := lastJudgeResult.Nominations[len(lastJudgeResult.Nominations)-1]

	lastNomination.Stages = append(lastNomination.Stages, &Stage{
		Title: line,
	})
}

func (fr *ForumResults) addPlace(line string) {
	lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]
	lastNomination := lastJudgeResult.Nominations[len(lastJudgeResult.Nominations)-1]
	lastStage := lastNomination.Stages[len(lastNomination.Stages)-1]

	lastStage.Places = append(lastStage.Places, parsePlace(line))
}

func (fr *ForumResults) addTechStage(line string) {
	lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]
	lastNomination := lastJudgeResult.Nominations[len(lastJudgeResult.Nominations)-1]

	lastNomination.TechStages = append(lastNomination.TechStages, &TechStage{
		Title: line,
	})
}

// 658   | AD|E   ==> выход в 1/4 финала ||      672   | BE|   место: 23-30
func (fr *ForumResults) addTechResult(line string) {
	lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]
	lastNomination := lastJudgeResult.Nominations[len(lastJudgeResult.Nominations)-1]
	lastTechStage := lastNomination.TechStages[len(lastNomination.TechStages)-1]

	lastTechStage.Results = append(lastTechStage.Results, &TechResult{
		Value: line,
	})
}

func (fr *ForumResults) addFinalTechStage(line string) {
	lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]
	lastNomination := lastJudgeResult.Nominations[len(lastJudgeResult.Nominations)-1]
	lastNomination.FinalTechStage = line
}

//  687   | 1 5 4 2 6         │     5  │ 4 4 3 2 5         │     4  │                   │        │    9 │    4
func (fr *ForumResults) addFinalTechResult(line string) {
	lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]
	lastNomination := lastJudgeResult.Nominations[len(lastJudgeResult.Nominations)-1]

	lastNomination.FinalTechResult = append(lastNomination.FinalTechResult, &FinalTechResult{
		Value: line,
	})
}
