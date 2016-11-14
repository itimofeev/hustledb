package prereg

import (
	"github.com/itimofeev/hustledb/util"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

func NewPreregService(db *runner.DB) *PreregService {
	return &PreregService{
		dao: NewPreregDao(db),
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

	for fCompId, url := range fCompUrls {
		if len(url) == 0 {
			continue
		}
		r := ParsePreregCompetition(fCompId, url)

		inserter := NewPreregInserter(util.GetDb())
		inserter.Insert(r)
	}
}
