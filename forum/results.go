package forum

type ForumResults struct {
	Judges []string
}

func (fr *ForumResults) addJudge(line string) {
	fr.Judges = append(fr.Judges, line)
}
