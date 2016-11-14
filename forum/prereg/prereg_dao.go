package prereg

import (
	"github.com/itimofeev/hustledb/forum/comp"
	"github.com/itimofeev/hustledb/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

func NewPreregDao(db *runner.DB, session *mgo.Session) *PreregDao {
	return &PreregDao{db, session}
}

type PreregDao struct {
	db      *runner.DB
	session *mgo.Session
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

func (d *PreregDao) InsertPreregComp(comp *PreregComp) *PreregComp {
	s := d.session.Clone()
	defer s.Close()

	c := s.DB("prereg_info").C("prereg_comp")

	found := make([]PreregComp, 0)
	err := c.Find(bson.M{"f_competition_id": comp.FCompetitionId}).All(&found)
	util.CheckErr(err)

	for _, foundComp := range found {
		err = c.RemoveId(foundComp.ID)
		util.CheckErr(err)
	}

	err = c.Insert(comp)
	util.CheckErr(err)
	return comp
}

func (d *PreregDao) ListPreregs() []PreregComp {
	s := d.session.Clone()
	defer s.Close()
	c := s.DB("prereg_info").C("prereg_comp")

	found := make([]PreregComp, 0)
	err := c.Find(bson.M{}).All(&found)
	util.CheckErr(err)

	return found
}

func (d *PreregDao) GetPreregById(fCompId int) *PreregComp {
	s := d.session.Clone()
	defer s.Close()
	c := s.DB("prereg_info").C("prereg_comp")

	found := &PreregComp{}
	err := c.Find(bson.M{"f_competition_id": fCompId}).One(&found)
	util.CheckErr(err)

	return found
}
