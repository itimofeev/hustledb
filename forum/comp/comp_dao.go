package comp

import (
	"github.com/itimofeev/hustledb/util"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"time"
)

var FCompetitionInsert = []string{"url", "title", "date", "description", "city", "approved_ash", "raw_text", "raw_text_changed", "download_date", "has_change"}

type FCompetition struct {
	ID             int64     `json:"id" db:"id"`
	Url            string    `json:"url" db:"url"`
	Title          string    `json:"title" db:"title"`
	Date           time.Time `json:"date" db:"date"`
	Desc           string    `json:"desc" db:"description"`
	City           *string   `json:"city" db:"city"`
	ApprovedASH    bool      `json:"approved" db:"approved_ash"` // Утверждено АСХ, конкурс не отменён, рейтинг не снят
	RawText        string    `json:"raw_text" db:"raw_text"`
	RawTextChanged string    `json:"raw_text_changed" db:"raw_text_changed"`
	DownloadDate   time.Time `json:"download_date" db:"download_date"`
	HasChange      bool      `db:"has_change"`
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

func (d *FCompDaoImpl) UpdateHasChange(existed* FCompetition, downloadDate time.Time, currentRaw string) {
	_, err := d.db.
		Update("f_competition").
		Set("raw_text_changed", currentRaw).
		Set("has_change", true).
		Set("download_date", downloadDate).
		Where("id = $1", existed.ID).
		Exec()

	util.CheckErr(err)
}

func (d *FCompDaoImpl) UpdateChangeFixed(existed* FCompetition) {
	_, err := d.db.
		Update("f_competition").
		Set("raw_text", existed.RawTextChanged).
		Set("has_change", false).
		Where("id = $1", existed.ID).
		Exec()

	util.CheckErr(err)
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
