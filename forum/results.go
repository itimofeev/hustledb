package forum

type ForumResults struct {
	JudgesResults []*JudgeResult

	Stages          []string `json:"-"`
	TechStages      []string `json:"-"`
	FinalTechStages []string `json:"-"`
	TechResult      []string `json:"-"`
	FinalTechResult []string `json:"-"`
	Places          []string `json:"-"`
}

type JudgeResult struct {
	Judges      []string
	Nominations []*NominationResult
}

type NominationResult struct {
	Title string
}

func (fr *ForumResults) addJudge(line string) {
	needNew := len(fr.JudgesResults) == 0 || len(fr.JudgesResults[len(fr.JudgesResults)-1].Nominations) > 0
	if needNew {
		fr.JudgesResults = append(fr.JudgesResults, &JudgeResult{
			Judges: []string{line},
		})
	} else {
		lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]
		lastJudgeResult.Judges = append(lastJudgeResult.Judges, line)
	}
}

func (fr *ForumResults) addNominationName(line string) {
	lastJudgeResult := fr.JudgesResults[len(fr.JudgesResults)-1]

	lastJudgeResult.Nominations = append(lastJudgeResult.Nominations, &NominationResult{
		Title: line,
	})
}

func (fr *ForumResults) addStage(line string) {
	fr.Stages = append(fr.Stages, line)
}

func (fr *ForumResults) addTechStage(line string) {
	fr.TechStages = append(fr.TechStages, line)
}

func (fr *ForumResults) addFinalTechStage(line string) {
	fr.FinalTechStages = append(fr.FinalTechStages, line)
}

func (fr *ForumResults) addTechResult(line string) {
	// 658   | AD|E   ==> выход в 1/4 финала ||      672   | BE|   место: 23-30
	fr.TechResult = append(fr.TechResult, line)
}

func (fr *ForumResults) addFinalTechResult(line string) {
	//  687   | 1 5 4 2 6         │     5  │ 4 4 3 2 5         │     4  │                   │        │    9 │    4
	fr.FinalTechResult = append(fr.FinalTechResult, line)
}

func (fr *ForumResults) addPlace(line string) {
	fr.Places = append(fr.Places, line)
}
