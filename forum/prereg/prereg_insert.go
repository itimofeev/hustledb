package prereg

import (
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

func NewPreregInserter(db *runner.DB) *PreregInserter {
	return &PreregInserter{dao: NewPreregDao(db)}
}

type PreregInserter struct {
	dao *PreregDao
}

func (pf *PreregInserter) Insert(c *PreregComp) {
	fComp := pf.dao.FindCompByForumUrl(c.FCompetitionUrl)
	c.FCompetitionId = fComp.ID

	pf.dao.InsertPreregComp(c)
}
