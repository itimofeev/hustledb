package hsaxls

import (
	"database/sql"
	"fmt"
	"github.com/itimofeev/hustledb/components/hsaxls/parser"
	"github.com/itimofeev/hustledb/components/util"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

const debug = false

func RepoListCompetitions(params util.PageParams) util.PageResponse {
	var competitions []parser.RawCompetition
	var total int

	sb := SqlBuilder{
		Select:  "*",
		From:    `competition c`,
		OrderBy: "c.date desc",
		Pp:      params,
	}

	pageQuery(util.DB, params, sb, &total, &competitions)

	return NewPageResponse(params, total, competitions)
}

func RepoGetCompetitionInfo(id int64) CompetitionDto {
	var comp CompetitionDto
	err := DoInTransaction(func(conn runner.Connection) error {
		rawComp := *GetCompetition(conn, id)
		rawNom := GetNominationsByCompetitionId(conn, id)

		comp.ID = rawComp.ID
		comp.Title = rawComp.Title

		var noms []NominationDto
		for _, rawNom := range *rawNom {
			noms = append(noms, NominationDto{
				ID:      rawNom.ID,
				Title:   rawNom.Value,
				Results: GetResultsByNomination(conn, rawNom.ID),
			})
		}

		comp.Nominations = noms

		return nil
	})
	if err != nil {
		panic(err)
	}

	return comp
}

func GetCompetition(conn runner.Connection, id int64) *parser.RawCompetition {
	var competition parser.RawCompetition
	err := conn.SQL(`
		SELECT
			*
		FROM
			competition c
		WHERE
			c.id = $1
	`, id).
		QueryStruct(&competition)

	if err == sql.ErrNoRows {
		panic(err)
	}

	return &competition
}

func GetResultsByNomination(conn runner.Connection, id int64) []NominationResultDto {
	var results []NominationResultDto

	/*
		ID int64 `json:"id" db:"id"`

		ResultString string `json:"resultString" db:"result"`

		DancerId int64 `json:"dancerId" db:"dancer_id"`
		DancerTitle string `json:"dancerTitle" db:"dancer_title"`


		IsJNJ bool `json:"isJnj"db:"is_jnj"`

		Points int    `json:"points" db:"points"`
		Class  string `json:"class" db:"class"`

		Place string `json:"place" db:"place"`
	*/

	err := conn.SQL(`
		SELECT
			r.id,
			r.result,
			r.is_jnj,
			r.points,
			r.class,
			r.all_places_from place,
			d.id dancer_id,
			d.surname || ' ' || d.name || coalesce(' ' || d.patronymic, '') dancer_title
		FROM
			result r
			join dancer d on r.dancer_id = d.id
		WHERE
			r.nomination_id = $1
		ORDER BY
			r.all_places_from asc
	`, id).QueryStructs(&results)
	if err != nil {
		panic(err)
	}

	return results
}

func RepoListDancers(params util.ListDancerParams) util.PageResponse {
	pageParams := util.PageParams{Limit: params.Limit, Offset: params.Offset}
	var dancers []parser.RawDancer

	sb := SqlBuilder{
		Select:  "*",
		From:    "dancer d",
		OrderBy: "d.id",
		Pp:      pageParams,
		Where:   "d.code ilike $1",
		Args:    []interface{}{fmt.Sprintf("%%%s%%", params.Query)},
	}

	var total int

	pageQuery(util.DB, pageParams, sb, &total, &dancers)

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

func GetDancerResults(conn runner.Connection, dancerId int64) *[]parser.RawCompetitionResult {
	var results []parser.RawCompetitionResult

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

func GetDancerClubs(conn runner.Connection, dancerId int64) *[]parser.RawClub {
	var clubs []parser.RawClub

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

func GetNominationsByCompetitionId(conn runner.Connection, compId int64) *[]parser.RawNomination {
	var nominations []parser.RawNomination

	err := conn.SQL(`
		SELECT
			*
		FROM
			nomination n
		WHERE
			n.competition_id = $1
		ORDER BY
			n.value asc
	`, compId).QueryStructs(&nominations)
	if err != nil {
		panic(err)
	}

	return &nominations
}

func GetDancer(conn runner.Connection, dancerId int64) *parser.RawDancer {
	var dancer parser.RawDancer
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
	Pp      util.PageParams
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

	sb.Pp.Fix()

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

func NewPageResponse(params util.PageParams, total int, slice interface{}) util.PageResponse {
	//TODO wtf!!!
	return util.PageResponse{Count: 20, Content: slice, PageSize: params.Limit, TotalCount: total}
}

func pageQuery(conn runner.Connection, params util.PageParams, sb SqlBuilder, total *int, result interface{}) util.PageResponse {
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
	tx, err := util.DB.Begin()
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
