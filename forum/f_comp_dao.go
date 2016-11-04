package forum

import (
	"github.com/itimofeev/hustledb/util"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"time"
)

type FCompDao interface {
	ListCompetitions() []FCompetition
	CreateCompetition(lat LinkAndTitle) *FCompetition
}

type FCompetition struct {
	ID    int64     `json:"id" db:"id"`
	Url   string    `json:"url" db:"url"`
	Title string    `json:"title" db:"title"`
	Date  time.Time `json:"date" db:"date"`
	Desc  string    `json:"desc" db:"description"`
	City  *string   `json:"city" db:"city"`
}

type FCompDaoImpl struct {
	db *runner.DB
}

func NewCompDao(db *runner.DB) FCompDao {
	return &FCompDaoImpl{
		db: db,
	}
}

func (d *FCompDaoImpl) CreateCompetition(lat LinkAndTitle) *FCompetition {
	comp := &FCompetition{
		Url:   lat.Link,
		Title: lat.Title,
		Desc:  lat.Desc,
		Date:  parseDate(lat.DateStr),
	}

	err := d.db.
		InsertInto("f_competition").
		Columns("url", "title", "date", "description").
		Record(comp).
		Returning("id").
		QueryScalar(&comp.ID)
	util.CheckErr(err)

	return comp
}

func parseDate(dateStr string) time.Time {
	year := dateStr[:4]
	monthStr := dateStr[5:7]
	day := dateStr[8:10]

	month := time.Month(util.Atoi(monthStr))

	return time.Date(util.Atoi(year), month, util.Atoi(day), 0, 0, 0, 0, time.UTC)
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
