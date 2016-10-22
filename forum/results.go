package forum

type ForumResults struct {
	JudgesResults []*Judge
}

type Judge struct {
	Judges      []string
	Nominations []*Nomination
}

type Nomination struct {
	Title           string
	Stages          []*Stage
	TechStages      []*TechStage
	FinalTechStage  string
	FinalTechResult []*FinalTechResult
}

type Places struct {
	Value string
}

type Stage struct {
	Title  string
	Places []*Places
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
	needNew := len(fr.JudgesResults) == 0 || len(fr.JudgesResults[len(fr.JudgesResults)-1].Nominations) > 0
	if needNew {
		fr.JudgesResults = append(fr.JudgesResults, &Judge{
			Judges: []string{line},
		})
	} else {
		lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]
		lastJudgeResult.Judges = append(lastJudgeResult.Judges, line)
	}
}

func (fr *ForumResults) addNominationName(line string) {
	lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]

	lastJudgeResult.Nominations = append(lastJudgeResult.Nominations, &Nomination{
		Title: line,
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

	lastStage.Places = append(lastStage.Places, &Places{
		Value: line,
	})
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
