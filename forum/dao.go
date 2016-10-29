package forum

import (
	"fmt"
	"github.com/iris-contrib/errors"
	"github.com/itimofeev/hustlesa/util"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"strings"
	"time"
)

func NewDao(db *runner.DB) ForumDao {
	return &ForumDaoImpl{
		db: db,
	}
}

type ForumDao interface {
	FindJudgeIdByTitle(title string) int64
	FindCompetitionId(time time.Time, title string) int64
}

type ForumDaoImpl struct {
	db *runner.DB
}

func (f *ForumDaoImpl) FindCompetitionId(time time.Time, title string) int64 {
	split := strings.Split(title, " ")
	var compIds []int64
	err := f.db.SQL(`
		SELECT
			c.id
		FROM
			competition c
		WHERE
			c.title ilike $1 || '%' AND c.date = $2
	`, split[0], time).
		QuerySlice(&compIds)

	util.CheckErr(err, fmt.Sprintf("Not found judge by title: %s", title))
	if len(compIds) != 1 {
		util.CheckErr(errors.New("WTF"), title, time)
	}

	return compIds[0]
}

func (f *ForumDaoImpl) FindJudgeIdByTitle(title string) int64 {
	if "Кудрявцева Ирина" == title { // TODO переделать на prevSurname
		return 1008
	}
	if "Графова Варвара" == title {
		return 1101
	}
	var dancerId int64
	err := f.db.SQL(`
		SELECT
			d.id
		FROM
			dancer d
		WHERE
			lower(d.surname || ' ' || d.name) = lower($1)
	`, strings.TrimSpace(title)).
		QueryScalar(&dancerId)

	util.CheckErr(err, fmt.Sprintf("Not found judge by title: %s", title))

	return dancerId
}
