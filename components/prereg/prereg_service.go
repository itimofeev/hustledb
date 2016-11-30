package prereg

import (
	"github.com/itimofeev/hustledb/components/util"
	"sync"
	"time"
)

var preregService *PreregService
var oncePreregService sync.Once

func initPreregService() {
	preregService = &PreregService{
		dao: &PreregDao{util.DB, util.MGO},
	}
}

func GetPreregService() *PreregService {
	oncePreregService.Do(initPreregService)
	return preregService
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

		s.insert(r)
	}
}

func (pf *PreregService) insert(c *PreregComp) {
	fComp := pf.dao.FindCompByForumUrl(c.FCompetitionUrl)
	c.FCompetitionId = fComp.ID
	c.UpdateDate = time.Now()

	pf.dao.insertPreregComp(c)
}
