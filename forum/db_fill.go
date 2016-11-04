package forum

func NewForumDbFiller(compUrl string, dao ForumDao) *ForumDbFiller {
	return &ForumDbFiller{
		dao:     dao,
		compUrl: compUrl,
	}
}

type ForumDbFiller struct {
	dao     ForumDao
	compUrl string
}

func (f *ForumDbFiller) FillDbInfo(fr *ForumResults) {
	fr.CompetitionId = f.dao.FindCompetitionId(f.compUrl)
	for _, judgeTeam := range fr.JudgesResults {
		f.fillJudges(fr.CompetitionId, judgeTeam.Judges)
	}
}

func (f *ForumDbFiller) fillJudges(compId int64, judges []*Judge) {
	for _, judge := range judges {
		if compId == 253 && judge.Title == "Барисова Ольга" { //Hustle & Discofox Festival Cup 2014
			judge.Title = "Борисова Ольга"
		}
		if compId == 253 && judge.Title == "Маликова Марика" { //Hustle & Discofox Festival Cup 2014
			judge.Title = "Маликова Мария"
		}
		dbJudgeId := f.dao.FindJudgeIdByTitle(judge.Title)
		judge.DancerId = dbJudgeId
	}
}
