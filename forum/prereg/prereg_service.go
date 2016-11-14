package prereg

import (
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
