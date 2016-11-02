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
	switch title {
	case "Лебедев Сергей", "Сергей Лебедев":
		return 952
	case "Кудрявцева Ирина":
		return 1008
	case "Графова Варвара":
		return 1101
	case "Потапов Вячеслав":
		return 1410
	case "Новикова Екатерина":
		return 3221
	case "Селиванова Алёна":
		return 2057
	}

	var dancerIds []int64
	err := f.db.SQL(`
		SELECT
			d.id
		FROM
			dancer d
		WHERE
			lower(d.surname || ' ' || d.name) = lower($1) OR
			lower(d.name || ' ' || d.surname) = lower($1)
	`, strings.TrimSpace(title)).
		QuerySlice(&dancerIds)

	util.CheckErr(err, fmt.Sprintf("Not found judge by title: %s", title))

	if len(dancerIds) != 1 {
		util.CheckOk(false, dancerIds, title)
	}

	return dancerIds[0]
}
