package forum

type ForumResults struct {
	Judges          []string
	NominationNames []string
	Stages          []string
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

func (fr *ForumResults) addPlace(line string) {
	fr.Places = append(fr.Places, line)
}
