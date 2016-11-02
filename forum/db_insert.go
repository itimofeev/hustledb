package forum

import (
	"database/sql"
	"fmt"
	"github.com/itimofeev/hustlesa/util"
	"github.com/labstack/gommon/log"
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

	in.insertDancerClubs(results)
}

func (i *DbInserter) clearRecordsForCompetitions(compId int64) {
	i.dao.DeleteDancerClubsByCompId(compId)
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
	nominationId := i.findNominationId(compId, n)
	n.RNominationId = NewNullInt64(nominationId)
	n = i.dao.CreateNomination(n)

	for _, stage := range n.Stages {
		stageTitle := parseStageTitle(stage)
		for _, place := range stage.Places {
			place.NominationId = n.ID
			place.StageTitle = stageTitle

			place.Dancer1Id = *i.findDancerId(compId, place.Dancer1)
			place.Result1Id = NewNullInt64(i.findResultId(compId, place.Dancer1Id, nominationId))
			dancer2Id := i.findDancerId(compId, place.Dancer2)
			if dancer2Id == nil {
				place.Dancer2Id = sql.NullInt64{Valid: false}
				place.Result2Id = sql.NullInt64{Valid: false}
			} else {
				place.Dancer2Id = sql.NullInt64{Valid: true, Int64: *dancer2Id}
				place.Result2Id = NewNullInt64(i.findResultId(compId, *dancer2Id, nominationId))
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

func (i *DbInserter) findNominationId(compId int64, n *Nomination) *int64 {
	for _, stage := range n.Stages {
		for _, somePlace := range stage.Places {
			dancerId := *i.findDancerId(compId, somePlace.Dancer1)
			isJnj := somePlace.Dancer2 == nil

			result := i.dao.FindResult(compId, dancerId, isJnj)

			if result != nil {
				return &result.NominationID
			}
		}
	}

	log.Warnf("Unable to find nomination_id for comp: %d and title %s", compId, n.Title)

	return nil
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

func (i *DbInserter) findResultId(compId, dancerId int64, nomId *int64) *int64 {
	if nomId == nil {
		return nil
	}
	return i.dao.FindResultNom(compId, *nomId, dancerId)
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

func (i *DbInserter) insertDancerClubs(fr *ForumResults) {
	dancerToClubs := make(map[int64]map[int64]bool)

	places := getAllPlaces(fr)

	for _, place := range places {
		i.addDancerClubs(dancerToClubs, place.Dancer1Id, place.Dancer1.Clubs)
		if place.Dancer2Id.Valid {
			i.addDancerClubs(dancerToClubs, place.Dancer2Id.Int64, place.Dancer2.Clubs)
		}
	}

	for dancerId, clubs := range dancerToClubs {
		for clubId := range clubs {
			i.dao.InsertDancerClub(fr.CompetitionId, dancerId, clubId)
		}
	}
}

func (i *DbInserter) addDancerClubs(dancerToClubs map[int64]map[int64]bool, dancerId int64, clubs []string) {
	if dancerId == 0 {
		log.Fatalf("Dancer %+v has no id", dancerId)
		return
	}

	if _, ok := dancerToClubs[dancerId]; ok {
		return
	}

	dancerClubs := make(map[int64]bool)
	for _, clubName := range clubs {
		clubId := i.dao.FindClubByName(clubName)
		if clubId == nil {
			log.Fatalf("Club %s not found", clubName)
		}
		dancerClubs[*clubId] = true
	}

	dancerToClubs[dancerId] = dancerClubs
}

func getAllPlaces(fr *ForumResults) []*Place {
	places := make([]*Place, 0)
	for _, jr := range fr.JudgesResults {
		for _, nom := range jr.Nominations {
			for _, stage := range nom.Stages {
				for _, place := range stage.Places {
					places = append(places, place)
				}
			}
		}
	}
	return places
}
