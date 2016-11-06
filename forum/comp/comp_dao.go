package comp

import (
	"github.com/itimofeev/hustledb/util"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"time"
)

var FCompetitionInsert = []string{"url", "title", "date", "description", "city", "approved_ash", "raw_text"}

type FCompetition struct {
	ID            int64     `json:"id" db:"id"`
	Url           string    `json:"url" db:"url"`
	Title         string    `json:"title" db:"title"`
	Date          time.Time `json:"date" db:"date"`
	Desc          string    `json:"desc" db:"description"`
	City          *string   `json:"city" db:"city"`
	ApprovedASH   bool      `json:"approved" db:"approved_ash"` // Утверждено АСХ, конкурс не отменён, рейтинг не снят
	RawText       string    `json:"raw_text" db:"raw_text"`
	AdminVerified bool      `json:"admin_verified" db:"admin_verified"` // Админ hustledb проверил изменения
}

type FCompDaoImpl struct {
	db *runner.DB
}

func NewCompDao(db *runner.DB) *FCompDaoImpl {
	return &FCompDaoImpl{
		db: db,
	}
}

func (d *FCompDaoImpl) CreateCompetition(comp *FCompetition) *FCompetition {
	err := d.db.
		InsertInto("f_competition").
		Columns(FCompetitionInsert...).
		Record(comp).
		Returning("id").
		QueryScalar(&comp.ID)
	util.CheckErr(err)

	return comp
}

func (d *FCompDaoImpl) ListCompetitions() []FCompetition {
	var comps = make([]FCompetition, 0)
	err := d.db.SQL(`
		SELECT
			c.*
		FROM
			f_competition c
		ORDER BY
			date desc
	`).
		QueryStructs(&comps)

	util.CheckErr(err)

	return comps
}

func (d *FCompDaoImpl) FindCompByUrl(compUrl string) *FCompetition {
	var comps = make([]FCompetition, 0)
	err := d.db.SQL(`
		SELECT
			c.*
		FROM
			f_competition c
		WHERE
			c.url = $1
	`, compUrl).
		QueryStructs(&comps)

	util.CheckErr(err)

	util.CheckOk(len(comps) <= 1, len(comps))

	if len(comps) == 0 {
		return nil
	}

	return &comps[0]
}
