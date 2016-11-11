package prereg

import (
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

func NewPreregFiller(db *runner.DB) *PreregFiller {
	return &PreregFiller{dao: NewPreregDao(db)}
}

type PreregFiller struct {
	dao *PreregDao
}

func (pf *PreregFiller) Fill(c *PreregComp) {
	fComp := pf.dao.FindCompByForumUrl(c.FCompetitionUrl)
	c.FCompetitionId = fComp.ID

	//for _, nom := range c.Nominations {
	//	for _, record := range nom.Records {
	//
	//	}
	//}
}
