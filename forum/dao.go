package forum

import (
	"errors"
	"fmt"
	"github.com/itimofeev/hustlesa/util"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"strings"
)

func NewDao(db *runner.DB) ForumDao {
	return &ForumDaoImpl{
		db: db,
	}
}

type ForumDao interface {
	FindJudgeIdByTitle(title string) int64
	FindCompetitionId(compUrl string) int64
}

type ForumDaoImpl struct {
	db *runner.DB
}

func (f *ForumDaoImpl) FindCompetitionId(compUrl string) int64 {
	var compIds []int64
	err := f.db.SQL(`
		SELECT
			c.id
		FROM
			competition c
		WHERE
			c.site = $1
	`, compUrl).
		QuerySlice(&compIds)

	util.CheckErr(err, fmt.Sprintf("Not found judge by title: %s", compUrl))
	if len(compIds) != 1 {
		util.CheckErr(errors.New("WTF"), compUrl)
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
