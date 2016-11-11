package prereg

import (
	"github.com/itimofeev/hustledb/forum/comp"
	"github.com/itimofeev/hustledb/util"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

func NewPreregDao(db *runner.DB) *PreregDao {
	return &PreregDao{db}
}

type PreregDao struct {
	db *runner.DB
}

func (d *PreregDao) FindCompByForumUrl(forumUrl string) *comp.FCompetition {
	var comps []comp.FCompetition
	err := d.db.SQL(`
		SELECT
			c.*
		FROM
			f_competition c
		WHERE
			c.url = $1
	`, forumUrl).
		QueryStructs(&comps)

	util.CheckErr(err)

	util.CheckOk(len(comps) == 1)

	return &comps[0]
}
