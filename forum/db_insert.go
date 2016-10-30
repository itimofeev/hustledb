package forum

import (
	"database/sql"
	"fmt"
	"github.com/itimofeev/hustlesa/util"
	"strings"
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
	i.dao.DeletePlacesByCompId(compId)
	i.dao.DeleteJudgesByCompId(compId)
	i.dao.DeleteNominationsByCompId(compId)
	i.dao.DeletePartitionsByCompId(compId)
}

func (i *DbInserter) insertPartition(index int, compId int64, jt *JudgeTeam) {
	partId := i.dao.CreatePartition(index, compId)

	for _, judge := range jt.Judges {
		judge.PartitionId = partId
		i.insertJudge(judge)
	}

	for _, nomination := range jt.Nominations {
		nomination.PartitionId = partId
		i.insertNomination(compId, nomination)
	}
}

func (i *DbInserter) insertJudge(j *Judge) {
	i.dao.CreateJudge(j)
}

func (i *DbInserter) insertNomination(compId int64, n *Nomination) *Nomination {
	n.RNominationId = i.findNominationId(compId, n)
	n = i.dao.CreateNomination(n)

	for _, stage := range n.Stages {
		stageTitle := parseStageTitle(stage)
		for _, place := range stage.Places {
			place.NominationId = n.ID
			place.StageTitle = stageTitle

			place.Dancer1Id = *i.findDancerId(compId, place.Dancer1)
			place.Result1Id = NewNullInt64(i.findResultId(compId, n.RNominationId, place.Dancer1Id))
			dancer2Id := i.findDancerId(compId, place.Dancer2)
			if dancer2Id == nil {
				place.Dancer2Id = sql.NullInt64{Valid: false}
				place.Result2Id = sql.NullInt64{Valid: false}
			} else {
				place.Dancer2Id = sql.NullInt64{Valid: true, Int64: *dancer2Id}
				place.Result2Id = NewNullInt64(i.findResultId(compId, n.RNominationId, *dancer2Id))
			}

			i.dao.CreatePlace(place)
		}
	}

	return n
}

func NewNullInt64(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{}
	} else {
		return sql.NullInt64{Valid: true, Int64: *i}
	}
}

func (i *DbInserter) findNominationId(compId int64, n *Nomination) int64 {
	somePlace := n.Stages[0].Places[0]
	dancerId := *i.findDancerId(compId, somePlace.Dancer1)
	isJnj := somePlace.Dancer2 == nil

	result := i.dao.FindResult(compId, dancerId, isJnj)

	fmt.Printf("For comp: %d, nom: %s found result: %s and nomination: %d\n", compId, n.Title, result.Result, result.NominationID)

	return result.NominationID
}

func (i *DbInserter) findDancerId(compId int64, dancer *Dancer) *int64 {
	if dancer == nil {
		return nil
	}

	if dancer.Id != 0 {
		id := int64(dancer.Id)
		return &id
	}

	dancerId := i.dao.FindDancer(compId, dancer.Title)

	if dancerId == nil {
		util.CheckOk(false, fmt.Sprintf("Unable to find dancer %+v in comp (%d)", dancer, compId))
	}

	return dancerId
}

func (i *DbInserter) findResultId(compId, nomId, dancerId int64) *int64 {
	return i.dao.FindResultNom(compId, nomId, dancerId)
}

func parseStageTitle(stage *Stage) string {
	switch {
	case stage.Title == "ФИНАЛ":
		return "1/1"
	case strings.Contains(stage.Title, "1/64 "):
		return "1/64"
	case strings.Contains(stage.Title, "1/32 "):
		return "1/32"
	case strings.Contains(stage.Title, "1/16 "):
		return "1/16"
	case strings.Contains(stage.Title, "1/8 "):
		return "1/8"
	case strings.Contains(stage.Title, "1/4 "):
		return "1/4"
	case strings.Contains(stage.Title, "1/2 "):
		return "1/2"
	default:
		util.CheckOk(false, "Unrecognized stage title: "+stage.Title)
	}

	return "error"
}
