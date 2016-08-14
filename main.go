package main

import (
	"bufio"
	"database/sql"
	"github.com/itimofeev/hustlesa/db/migrations"
	_ "github.com/lib/pq" //postgres driver
	"github.com/rubenv/sql-migrate"
	"gopkg.in/mgutz/dat.v1"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"log"
	"os"
	"strings"
	"time"

	"fmt"
	"github.com/itimofeev/hustlesa/model"
	"github.com/itimofeev/hustlesa/parser"
	"reflect"
)

func initDb(config Config) *runner.DB {
	db, err := sql.Open("postgres", config.Db().URL)

	CheckErr(err, "Open db")

	runner.MustPing(db)

	db.SetMaxIdleConns(config.Db().MaxIdleConns)
	db.SetMaxOpenConns(config.Db().MaxOpenConns)

	// migrations
	migrations := &migrate.AssetMigrationSource{
		Asset:    migrations.Asset,
		AssetDir: migrations.AssetDir,
		Dir:      "db/migrations",
	}

	_, err = migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Panic("Migrations not applied, err " + err.Error())
	}

	// set this to enable interpolation
	dat.EnableInterpolation = true

	// set to check things like sessions closing.
	// Disable strick mode for production
	dat.Strict = config.Db().StrictMode

	// Log any query over 10ms as warnings. (optional)
	runner.LogQueriesThreshold = 10 * time.Millisecond

	return runner.NewDB(db, "postgres")
}

//Устанавливает переменные окружения, заданные в testing.env.list
func initEnvironment() {
	file, err := os.Open("/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustlesa/tools/local.env")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.SplitN(scanner.Text(), "=", 2)
		if len(s) < 2 {
			//skip empty lines
			continue
		}

		os.Setenv(s[0], s[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initEnvironment()

	config := ReadConfig()

	db := initDb(config)

	fmt.Println("!!! ", db) //TODO remove

	res := parser.Parse("/Users/ilyatimofee/prog/hsa/parse-xls/json/")

	fmt.Printf("!!!%+v\n", (*res.Dancers)[10]) //TODO remove

	//clubs := parser.ParseClubs("/Users/ilyatimofee/prog/hsa/parse-xls/json/clubs.json")
	//
	//for _, club := range *clubs {
	//	fixString(&(*clubs)[0])
	//	_, err := insertClub(db, &club)
	//
	//	CheckErr(err, "insert club")
	//}
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

func insertClub(db *runner.DB, club *model.RawClub) (*model.RawClub, error) {
	err := db.
		InsertInto("club").
		Columns("id", "name", "leader", "comment", "site1", "old_name").
		Record(club).
		Returning("id").
		QueryScalar(&club.ID)

	return club, err
}
