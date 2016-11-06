package comp

import (
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

func NewFCompService(db *runner.DB) *FCompService {
	return &FCompService{
		dao: NewCompDao(db),
	}
}

type FCompService struct {
	dao *FCompDaoImpl
}

func (s *FCompService) ListCompetitions() []FCompetition {
	return s.dao.ListCompetitions()
}

func (s *FCompService) ParseCompetitions() {
	comps := ParseCompetitionListFromForum()
	for _, comp := range comps {
		s.ProcessCompetition(&comp)
	}
}

func (s *FCompService) ProcessCompetition(fromForum *FCompetition) {
	inDb := s.dao.FindCompByUrl(fromForum.Url)
	if inDb == nil {
		s.dao.CreateCompetition(fromForum)
	}
}
