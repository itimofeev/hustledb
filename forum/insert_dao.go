package forum

import (
	"github.com/itimofeev/hustlesa/util"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

func NewInsertDao(db *runner.DB) InsertDao {
	return &InsertDaoImpl{
		db: db,
	}
}

type InsertDao interface {
	CreatePartition(index int, compId int64) int64
	DeletePartitionsByCompId(compId int64)
}

type InsertDaoImpl struct {
	db *runner.DB
}

func (d *InsertDaoImpl) DeletePartitionsByCompId(compId int64) {
	_, err := d.db.
		DeleteFrom("partition").
		Where("competition_id = $1", compId).
		Exec()
	util.CheckErr(err)
}

func (d *InsertDaoImpl) CreatePartition(index int, compId int64) int64 {
	partition := struct {
		ID     int64 `db:"id"`
		Index  int   `db:"index"`
		CompId int64 `db:"competition_id"`
	}{
		Index:  index,
		CompId: compId,
	}
	err := d.db.
		InsertInto("partition").
		Columns("index", "competition_id").
		Record(partition).
		Returning("id").
		QueryScalar(&partition.ID)
	util.CheckErr(err)

	return partition.ID
}
