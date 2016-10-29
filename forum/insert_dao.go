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
	DeleteJudgesByCompId(compId int64)
	CreateJudge(j *Judge) *Judge
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

func (d *InsertDaoImpl) DeleteJudgesByCompId(compId int64) {
	_, err := d.db.
		DeleteFrom("judge j").
		Where("exists (SELECT NULL FROM partition p WHERE p.competition_id = $1 AND j.partition_id = p.id)", compId).
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

func (d *InsertDaoImpl) CreateJudge(judge *Judge) *Judge {
	err := d.db.
		InsertInto("judge").
		Columns("letter", "partition_id", "dancer_id").
		Record(judge).
		Returning("id").
		QueryScalar(&judge.ID)
	util.CheckErr(err)

	return judge
}
