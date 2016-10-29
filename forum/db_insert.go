package forum

import (
	"fmt"
)

func NewDbInserter(dao InsertDao) *DbInserter {
	return &DbInserter{
		dao: dao,
	}
}

type DbInserter struct {
	dao InsertDao
}

func (in *DbInserter) Insert(results *ForumResults) {
	in.clearRecordsForCompetitions(results.CompetitionId)
	for i, jr := range results.JudgesResults {
		in.insertPartition(i+1, results.CompetitionId, jr)
	}
}

func (i *DbInserter) clearRecordsForCompetitions(compId int64) {
	i.dao.DeletePartitionsByCompId(compId)
}

func (i *DbInserter) insertPartition(index int, compId int64, jt *JudgeTeam) {
	partId := i.dao.CreatePartition(index, compId)
	fmt.Println("!!!", partId) //TODO remove
}
