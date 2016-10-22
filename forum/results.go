package forum

type ForumResults struct {
	Judges          []string
	NominationNames []string
	Stages          []string
	TechStages      []string
	FinalTechStages []string
	TechResult      []string
	FinalTechResult []string
	Places          []string
}

func (fr *ForumResults) addJudge(line string) {
	fr.Judges = append(fr.Judges, line)
}

func (fr *ForumResults) addNominationName(line string) {
	fr.NominationNames = append(fr.NominationNames, line)
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
