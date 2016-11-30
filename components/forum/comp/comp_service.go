package comp

import (
	"github.com/itimofeev/hustledb/components/util"
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
	comps, err := ParseCompetitionListFromForum()
	if err != nil {
		util.CompLog.WithError(err).Error("Unable to parse competition list from forum")
	}
	for _, comp := range comps {
		s.ProcessCompetition(&comp)
	}
}

func (s *FCompService) ProcessCompetition(fromForum *FCompetition) {
	inDb := s.dao.FindCompByUrl(fromForum.Url)
	// в базе такой записи вообще нет, создаём
	if inDb == nil {
		s.dao.CreateCompetition(fromForum)
		return
	}
	// запись в базе есть
	if inDb.RawText == fromForum.RawText {
		// ничего не поменялось с даты последнего обновления -> выходим
		return
	}
	// запись об изменении уже есть
	if inDb.RawTextChanged == fromForum.RawText {
		return

	}
	s.dao.UpdateHasChange(inDb, fromForum.DownloadDate, fromForum.RawText)
}
