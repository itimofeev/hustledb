package forum

import (
	"fmt"
	"github.com/itimofeev/hustlesa/util"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"strings"
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
	DeletePlacesByCompId(compId int64)
	DeleteNominationsByCompId(compId int64)
	CreateJudge(j *Judge) *Judge
	CreateNomination(n *Nomination) *Nomination
	CreatePlace(p *Place) *Place
	FindDancer(compId int64, dTitle string) *int64
}

type InsertDaoImpl struct {
	db *runner.DB
}

func (d *InsertDaoImpl) DeletePartitionsByCompId(compId int64) {
	_, err := d.db.
		DeleteFrom("f_partition").
		Where("competition_id = $1", compId).
		Exec()
	util.CheckErr(err)
}

func (d *InsertDaoImpl) DeleteJudgesByCompId(compId int64) {
	_, err := d.db.
		DeleteFrom("f_judge j").
		Where("exists (SELECT NULL FROM f_partition p WHERE p.competition_id = $1 AND j.partition_id = p.id)", compId).
		Exec()
	util.CheckErr(err)
}

func (d *InsertDaoImpl) DeleteNominationsByCompId(compId int64) {
	_, err := d.db.
		DeleteFrom("f_nomination n").
		Where("exists (SELECT NULL FROM f_partition p WHERE p.competition_id = $1 AND n.partition_id = p.id)", compId).
		Exec()
	util.CheckErr(err)
}

func (d *InsertDaoImpl) DeletePlacesByCompId(compId int64) {
	_, err := d.db.
		DeleteFrom("f_place pl").
		Where("exists ("+
			"       SELECT "+
			"               NULL "+
			"       FROM "+
			"               f_partition p "+
			"               JOIN f_nomination n on p.id = n.partition_id"+
			"       WHERE "+
			"               p.competition_id = $1 AND pl.nomination_id = n.id"+
			"       )", compId).
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
		InsertInto("f_partition").
		Columns("index", "competition_id").
		Record(partition).
		Returning("id").
		QueryScalar(&partition.ID)
	util.CheckErr(err)

	return partition.ID
}

func (d *InsertDaoImpl) CreateJudge(judge *Judge) *Judge {
	err := d.db.
		InsertInto("f_judge").
		Columns("letter", "partition_id", "dancer_id").
		Record(judge).
		Returning("id").
		QueryScalar(&judge.ID)
	util.CheckErr(err)

	return judge
}

func (d *InsertDaoImpl) CreateNomination(nomination *Nomination) *Nomination {
	err := d.db.
		InsertInto("f_nomination").
		Columns("title", "partition_id").
		Record(nomination).
		Returning("id").
		QueryScalar(&nomination.ID)
	util.CheckErr(err)

	return nomination
}

func (d *InsertDaoImpl) FindDancer(compId int64, dTitle string) *int64 {
	dTitle2 := strings.Replace(dTitle, "ё", "е", -1)
	var dancerId int64
	err := d.db.SQL(`
		SELECT
			d.id
		FROM
			dancer d
		WHERE
			($1 ilike '%' || d.name || '%' OR $2 ilike '%' || d.name || '%') AND
			($1 ilike d.surname || '%' OR $2 ilike d.surname || '%' OR $1 ilike d.prev_surname || '%')
	`, dTitle, dTitle2).
		QueryScalar(&dancerId)

	util.CheckErr(err, dTitle)

	return &dancerId
}

func (d *InsertDaoImpl) CreatePlace(p *Place) *Place {
	err := d.db.
		InsertInto("f_place").
		Columns("place_from", "place_to", "number", "stage_title", "nomination_id", "dancer1_id", "dancer2_id").
		Record(p).
		Returning("id").
		QueryScalar(&p.ID)
	util.CheckErr(err, fmt.Sprintf("%+v", p))

	return p
}
