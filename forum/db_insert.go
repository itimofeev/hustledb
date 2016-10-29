package forum

import "gopkg.in/mgutz/dat.v1/sqlx-runner"

func NewDbInserter(db *runner.DB) *DbInserter {
	return &DbInserter{
		db: db,
	}
}

type DbInserter struct {
	db *runner.DB
}

func (i *DbInserter) Insert(results *ForumResults) {

}
