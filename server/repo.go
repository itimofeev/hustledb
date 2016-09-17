package server

import (
	"database/sql"
	"fmt"
	"github.com/itimofeev/hustlesa/model"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

const debug = false

func RepoListCompetitions(params PageParams) PageResponse {
	var competitions []model.RawCompetition
	var total int

	sb := SqlBuilder{
		Select:  "*",
		From:    `competition c`,
		OrderBy: "c.id",
		Pp:      params,
	}

	pageQuery(db, params, sb, &total, &competitions)

	return NewPageResponse(params, total, competitions)
}

func RepoListDancers(params ListDancerParams) PageResponse {
	pageParams := PageParams{Limit: params.Limit, Offset: params.Offset}
	var dancers []model.RawDancer

	sb := SqlBuilder{
		Select:  "*",
		From:    "dancer d",
		OrderBy: "d.id",
		Pp:      pageParams,
		Where:   "d.code ilike $1",
		Args:    []interface{}{fmt.Sprintf("%%%s%%", params.Query)},
	}

	var total int

	pageQuery(db, pageParams, sb, &total, &dancers)

	return NewPageResponse(pageParams, total, dancers)
}

func RepoGetDancerInfo(dancerId int64) DancerProfileDto {
	var profile DancerProfileDto
	err := DoInTransaction(func(conn runner.Connection) error {
		profile = *GetDancerProfile(conn, dancerId)
		profile.Clubs = *GetDancerClubsDto(conn, dancerId)
		profile.Results = *GetDancerResultsDto(conn, dancerId)

		return nil
	})
	if err != nil {
		panic(err)
	}

	return profile
}

func GetDancerResults(conn runner.Connection, dancerId int64) *[]model.RawCompetitionResult {
	var results []model.RawCompetitionResult

	err := conn.SQL(`
		SELECT
			r.*
		FROM
			result r
			JOIN competition c ON r.competition_id = c.id
		WHERE
			r.dancer_id = $1
		ORDER BY
			c.date desc
	`, dancerId).QueryStructs(&results)
	if err != nil {
		panic(err)
	}

	return &results
}

func GetDancerResultsDto(conn runner.Connection, dancerId int64) *[]ResultDto {
	var results []ResultDto

	err := conn.SQL(`
		SELECT
			r.id,
			r.result,
			r.is_jnj,
			r.points,
			r.class,
			r.all_places_from place,

			n.id nom_id,
			n.value nom_title,

			c.id comp_id,
			c.title comp_title,
			c.date comp_date
		FROM
			result r
			JOIN competition c ON r.competition_id = c.id
			JOIN nomination n on r.nomination_id = n.id
		WHERE
			r.dancer_id = $1
		ORDER BY
			c.date desc
	`, dancerId).QueryStructs(&results)
	if err != nil {
		panic(err)
	}

	return &results
}

func GetDancerClubs(conn runner.Connection, dancerId int64) *[]model.RawClub {
	var clubs []model.RawClub

	err := conn.SQL(`
		SELECT
			c.*
		FROM
			club c
			JOIN dancer_club dc on c.id = dc.club_id
		WHERE
			dc.dancer_id = $1
		ORDER BY
			c.name asc
	`, dancerId).QueryStructs(&clubs)
	if err != nil {
		panic(err)
	}

	return &clubs
}

func GetDancerClubsDto(conn runner.Connection, dancerId int64) *[]ClubDto {
	var clubs []ClubDto

	err := conn.SQL(`
		SELECT
			c.id,
			c.name title
		FROM
			club c
			JOIN dancer_club dc on c.id = dc.club_id
		WHERE
			dc.dancer_id = $1
		ORDER BY
			c.name asc
	`, dancerId).QueryStructs(&clubs)
	if err != nil {
		panic(err)
	}

	return &clubs
}

func GetDancer(conn runner.Connection, dancerId int64) *model.RawDancer {
	var dancer model.RawDancer
	err := conn.SQL(`
		SELECT
			*
		FROM
			dancer d
		WHERE
			d.id = $1
	`, dancerId).
		QueryStruct(&dancer)

	if err == sql.ErrNoRows {
		panic(err)
	}

	return &dancer
}

func GetDancerProfile(conn runner.Connection, dancerId int64) *DancerProfileDto {

	var dancerProfile DancerProfileDto
	err := conn.SQL(`
		SELECT
			d.id,
			d.code,
			d.surname || ' ' || d.name || coalesce(' ' || d.surname, '') title,
			d.pair_class classic_class,
			d.jnj_class
		FROM
			dancer d
		WHERE
			d.id = $1
	`, dancerId).
		QueryStruct(&dancerProfile)

	if err == sql.ErrNoRows {
		panic(err)
	}

	return &dancerProfile
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

	sb.Pp.fix()

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

func NewPageResponse(params PageParams, total int, slice interface{}) PageResponse {
	//TODO wtf!!!
	return PageResponse{Count: 20, Content: slice, PageSize: params.Limit, TotalCount: total}
}

func pageQuery(conn runner.Connection, params PageParams, sb SqlBuilder, total *int, result interface{}) PageResponse {
	totalSql := sb.totalQuery()

	if debug {
		fmt.Println("Args: ", sb.Args)
		fmt.Println("TotalSql: ", totalSql)
	}

	if err := conn.SQL(totalSql, sb.Args...).QueryScalar(total); err != nil {
		panic(err)
	}

	if *total == 0 {
		return NewPageResponse(params, *total, result)
	}

	dataSql, args := sb.dataQuery()

	if debug {
		fmt.Println("Args: ", args)
		fmt.Println("DataSql: ", dataSql)
	}

	if err := conn.SQL(dataSql, args...).QueryStructs(result); err != nil {
		panic(err)
	}

	return NewPageResponse(params, *total, result)
}

// DoInTransaction executes function passed as parameter in single transaction
func DoInTransaction(atomicAction func(conn runner.Connection) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.AutoRollback()

	err = atomicAction(tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
