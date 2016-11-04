package parser

import (
	"fmt"
	"github.com/itimofeev/hustledb/model"
	"github.com/itimofeev/hustledb/util"
	"gopkg.in/mgutz/dat.v1"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"reflect"
	"strings"
)

func InsertData(db *runner.DB, res model.RawParsingResults) {
	for _, club := range res.Clubs {
		fixString(&club)
		_, err := insertClub(db, &club)

		util.CheckErr(err, "insert club")
	}

	for _, dancer := range res.Dancers {
		fixString(&dancer)
		_, err := insertDancer(db, &dancer)

		util.CheckErr(err, fmt.Sprintf("insert dancer: %+v", dancer))
	}

	for _, dc := range res.DancerClubs {
		fixString(&dc)

		_, err := insertDancerClub(db, &dc)

		util.CheckErr(err, "insert dancerClub")
	}

	for _, c := range res.Competitions {
		fixString(&c)

		_, err := insertCompetition(db, &c)

		util.CheckErr(err, "insert competition")
	}

	for _, n := range res.Nominations {
		fixString(&n)

		_, err := insertNomination(db, &n)

		util.CheckErr(err, "insert nomination")
	}

	for _, cr := range res.CompResults {
		fixString(&cr)

		_, err := insertCompResult(db, &cr)

		util.CheckErr(err, "insert comp results")
	}
}

func insertClub(db *runner.DB, club *model.RawClub) (*model.RawClub, error) {
	err := db.
		InsertInto("club").
		Columns("id", "name", "leader", "comment", "site1", "old_name").
		Record(club).
		Returning("id").
		QueryScalar(&club.ID)

	return club, err
}

func insertDancerClub(db *runner.DB, dancerClub *model.RawDancerClub) (*model.RawDancerClub, error) {
	_, err := db.
		InsertInto("dancer_club").
		Columns("dancer_id", "club_id").
		Record(dancerClub).
		Exec()

	return dancerClub, err
}

func insertDancer(db *runner.DB, dancer *model.RawDancer) (*model.RawDancer, error) {
	err := db.
		InsertInto("dancer").
		Columns("id", "code", "name", "surname", "patronymic", "sex", "pair_class", "jnj_class", "prev_surname", "source").
		Record(dancer).
		Returning("id").
		QueryScalar(&dancer.ID)

	return dancer, err
}

func insertCompetition(db *runner.DB, competition *model.RawCompetition) (*model.RawCompetition, error) {
	err := db.
		InsertInto("competition").
		Columns("id", "title", "date", "site").
		Record(competition).
		Returning("id").
		QueryScalar(&competition.ID)

	return competition, err
}

func insertNomination(db *runner.DB, nominations *model.RawNomination) (*model.RawNomination, error) {
	err := db.
		InsertInto("nomination").
		Columns("competition_id", "value", "male_count", "female_count", "type", "min_class", "max_class").
		Record(nominations).
		Returning("id").
		QueryScalar(&nominations.ID)

	return nominations, err
}

func insertCompResult(db *runner.DB, result *model.RawCompetitionResult) (*model.RawCompetitionResult, error) {
	err := db.
		InsertInto("result").
		Columns("competition_id", "dancer_id", "nomination_id", "result",
			"place", "place_from", "is_jnj", "points", "class",
			"all_places_from", "all_places_to", "all_places_min_class", "all_places_max_class").
		Record(result).
		Returning("id").
		QueryScalar(&result.ID)

	return result, err
}

func fixString(obj interface{}) {
	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		switch field.Interface().(type) {
		case string:
			field.SetString(strings.TrimSpace(field.String()))
		case dat.NullString:
			strField := field.FieldByName("String")
			validField := field.FieldByName("Valid")

			str := strField.String()
			valid := validField.Bool()
			if valid {
				str := strings.TrimSpace(str)
				strField.SetString(str)

				if "" == str {
					validField.SetBool(false)
				}
			}
		}
	}
}
