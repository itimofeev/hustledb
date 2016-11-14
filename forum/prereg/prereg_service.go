package prereg

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

func NewPreregService(db *runner.DB, session *mgo.Session) *PreregService {
	return &PreregService{
		dao: NewPreregDao(db, session),
	}
}

type PreregService struct {
	dao *PreregDao
}

func (s *PreregService) ListPreregs() []PreregComp {
	return s.dao.ListPreregs()
}

func (s *PreregService) GetPreregById(id int) *PreregComp {
	return s.dao.GetPreregById(id)
}

func (s *PreregService) ParsePreregInfo() {
	listLinks := ParseAllPreregLinks()

	var ids []int
	for _, listLink := range listLinks {
		ids = append(ids, ParsePreregId(listLink))
	}

	fCompUrls := make(map[int]string)
	for _, preregId := range ids {
		fCompUrls[preregId] = GetForumCompetitionId(preregId)
	}

	inserter := NewPreregInserter(s.dao.db, s.dao.session)

	for fCompId, url := range fCompUrls {
		if len(url) == 0 {
			continue
		}
		r := ParsePreregCompetition(fCompId, url)

		inserter.Insert(r)
	}
}
