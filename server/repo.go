package server

import (
	"fmt"
	"github.com/itimofeev/hustlesa/model"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"strings"
)

func RepoListCompetitions(params PageParams) PageResponse {
	var competitions []model.RawCompetition

	sql := `
		SELECT
			{COLUMNS}
		FROM
			competition c
	`
	var total int

	pageQuery(db, params, sql, &total, &competitions)

	return NewPageResponse(params, total, competitions)
}
func NewPageResponse(params PageParams, total int, slice interface{}) PageResponse {
	//TODO wtf!!!
	return PageResponse{Count: 20, Content: slice, PageSize: params.Limit, TotalCount: total}
}

func pageQuery(conn runner.Connection, params PageParams, sqlStr string, total *int, result interface{}, args ...interface{}) PageResponse {
	countSql := strings.Replace(sqlStr, "{COLUMNS}", "count(*)", -1)

	if err := conn.SQL(countSql, args...).QueryScalar(total); err != nil {
		panic(err)
	}

	if *total == 0 {
		return NewPageResponse(params, *total, result)
	}

	selectSql := strings.Replace(sqlStr, "{COLUMNS}", "*", -1)

	argsLen := len(args)
	selectSql = fmt.Sprintf("%s \n LIMIT $%d OFFSET $%d", selectSql, argsLen+1, argsLen+2)

	var fullArgs []interface{}
	fullArgs = append(fullArgs, args...)
	fullArgs = append(fullArgs, params.Limit, params.Offset)

	if err := conn.SQL(selectSql, fullArgs...).QueryStructs(result); err != nil {
		panic(err)
	}

	return NewPageResponse(params, *total, result)
}
