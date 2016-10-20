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

	"github.com/itimofeev/hustlesa/server"
	"github.com/itimofeev/hustlesa/util"
	"net/http"
)

func initDb(config util.Config) *runner.DB {
	db, err := sql.Open("postgres", config.Db().URL)

	util.CheckErr(err, "Open db")

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
	//initEnvironment()

	config := util.ReadConfig()

	db := initDb(config)

	//res := parser.Parse("config.App().JsonFilesPath")
	//parser.InsertData(db, res)

	log.Fatal(http.ListenAndServe(":8080", server.InitRouter(db)))

}
