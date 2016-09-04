package server

import (
	"fmt"
	"github.com/itimofeev/hustlesa/model"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

func RepoListCompetitions(params PageParams) PageResponse {
	var competitions []model.RawCompetition
	var total int

	sb := SqlBuilder{
		Select: "*",
		From: `FROM
			competition c`,
		OrderBy: "c.id",
		Pp:      params,
	}

	pageQuery(db, params, sb, &total, &competitions)

	return NewPageResponse(params, total, competitions)
}

type SqlBuilder struct {
	Select  string
	From    string
	OrderBy string
	Where   string
	Args    []interface{}
	Pp      PageParams
}

func (sb SqlBuilder) totalQuery() string {
	where := sb.Where
	if where == "" {
		where = "true"
	}
	return fmt.Sprintf(`
		SELECT
			count(*)
		FROM
			%s
		WHERE
			%s
	`, sb.From, where)
}

func (sb SqlBuilder) dataQuery() (string, []interface{}) {
	where := sb.Where
	if where == "" {
		where = "true"
	}
	orderBy := sb.OrderBy
	if orderBy != "" {
		orderBy = fmt.Sprintf(
			`ORDER BY
				%s`, orderBy,
		)
	}

	var fullArgs []interface{}
	fullArgs = append(fullArgs, sb.Args...)
	fullArgs = append(fullArgs, sb.Pp.Limit, sb.Pp.Offset)

	argsLen := len(sb.Args)

	return fmt.Sprintf(`
		SELECT
			%s
		FROM
			%s
		WHERE
			%s
		%s
		LIMIT	$%d
		OFFSET	$%d`, sb.Select, sb.From, where, orderBy, argsLen+1, argsLen+2,
		),
		fullArgs
}

func RepoListDancers(params PageParams) PageResponse {
	var dancers []model.RawDancer

	sb := SqlBuilder{
		Select:  "*",
		From:    "dancer d",
		OrderBy: "d.id",
		Pp:      params,
	}

	var total int

	pageQuery(db, params, sb, &total, &dancers)

	return NewPageResponse(params, total, dancers)
}

func NewPageResponse(params PageParams, total int, slice interface{}) PageResponse {
	//TODO wtf!!!
	return PageResponse{Count: 20, Content: slice, PageSize: params.Limit, TotalCount: total}
}

func pageQuery(conn runner.Connection, params PageParams, sb SqlBuilder, total *int, result interface{}) PageResponse {
	totalSql := sb.totalQuery()

	if err := conn.SQL(totalSql, sb.Args...).QueryScalar(total); err != nil {
		panic(err)
	}

	if *total == 0 {
		return NewPageResponse(params, *total, result)
	}

	dataSql, args := sb.dataQuery()
	fmt.Println("!!! ", dataSql) //TODO remove
	fmt.Println("!!! ", args)    //TODO remove
	if err := conn.SQL(dataSql, args...).QueryStructs(result); err != nil {
		panic(err)
	}

	return NewPageResponse(params, *total, result)
}
