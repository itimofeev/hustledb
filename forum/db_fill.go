package forum

func NewForumDbFiller(dao ForumDao) *ForumDbFiller {
	return &ForumDbFiller{
		dao: dao,
	}
}

type ForumDbFiller struct {
	dao ForumDao
}

func (f *ForumDbFiller) FillDbInfo(fr *ForumResults) {
	fr.CompetitionId = f.dao.FindCompetitionId(fr.Date, fr.Title)
	for _, judgeTeam := range fr.JudgesResults {
		f.fillJudges(judgeTeam.Judges)
	}
}

func (f *ForumDbFiller) fillJudges(judges []*Judge) {
	for _, judge := range judges {
		dbJudgeId := f.dao.FindJudgeIdByTitle(judge.Title)
		judge.DancerId = dbJudgeId
	}
}
