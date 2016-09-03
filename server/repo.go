package server

import "github.com/itimofeev/hustlesa/model"

func RepoListCompetitions(params PageParams) ([]model.RawCompetition, error) {
	var competitions []model.RawCompetition
	err := db.SQL(`
		SELECT
			*
		FROM
			competition
		ORDER BY
			id desc
		LIMIT $1
		OFFSET $2
	`,
		params.Limit,
		params.Offset,
	).QueryStructs(&competitions)
	return competitions, err
}
