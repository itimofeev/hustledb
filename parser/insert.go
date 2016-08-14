package parser

import (
	"github.com/itimofeev/hustlesa/model"
	"gopkg.in/mgutz/dat.v1"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"reflect"
	"strings"
)

func InsertData(db *runner.DB, res model.RawParsingResults) {
	for _, club := range res.Clubs {
		fixString(&club)
		_, err := insertClub(db, &club)

		CheckErr(err, "insert club")
	}

	for _, dancer := range res.Dancers {
		fixString(&dancer)
		_, err := insertDancer(db, &dancer)

		CheckErr(err, "insert club")
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

func insertDancer(db *runner.DB, dancer *model.RawDancer) (*model.RawDancer, error) {
	err := db.
		InsertInto("dancer").
		Columns("id", "code", "name", "surname", "patronymic", "sex", "pair_class", "jnj_class", "prev_surname", "source").
		Record(dancer).
		Returning("id").
		QueryScalar(&dancer.ID)

	return dancer, err
}

/***
  id BIGSERIAL PRIMARY KEY NOT NULL,

  code VARCHAR(256) NOT NULL,--TODO make unique

  name VARCHAR (256) NOT NULL,
  surname VARCHAR (256) NOT NULL,
  patronymic VARCHAR (256),
  sex VARCHAR(256) NOT NULL,

  pair_class VARCHAR(256) NOT NULL,
  jnj_class VARCHAR(256) NOT NULL,

  prev_surname VARCHAR(256),
  source VARCHAR(256)
*/

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
