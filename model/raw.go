package model

import (
	"gopkg.in/mgutz/dat.v1"
	"time"
)

type RawParsingResults struct {
	Clubs        []RawClub
	Dancers      []RawDancer
	DancerClubs  []RawDancerClub
	Competitions []RawCompetition
	Nominations  []RawNomination
}

type RawClub struct {
	ID      int64          `json:"id" db:"id"`
	Name    string         `json:"name" db:"name"`
	Leader  dat.NullString `json:"leader" db:"leader"`
	Site1   dat.NullString `json:"site1" db:"site1"`
	OldName dat.NullString `json:"oldName" db:"old_name"`
	Comment dat.NullString `json:"comment" db:"comment"`
}

type RawDancer struct {
	ID          int64          `json:"id" db:"id"`
	Title       string         `json:"name"`
	PrevSurname dat.NullString `json:"prevSurname" db:"prev_surname"`
	PairClass   string         `json:"pairClass" db:"pair_class"`
	JnjClass    string         `json:"jnjClass" db:"jnj_class"`
	Sex         string         `json:"sex" db:"sex"`
	Source      string         `json:"source" db:"source"`

	Code       string         `db:"code"`
	Name       string         `db:"name"`
	Surname    string         `db:"surname"`
	Patronymic dat.NullString `db:"patronymic"`
}

type RawDancerClub struct {
	DancerID  int64  `json:"dancerId" db:"dancer_id"`
	ClubNames string `json:"clubId"`
	ClubId    int64  `db:"club_id"`
}

type RawCompetition struct {
	ID      int64          `json:"id" db:"id"`
	Title   string         `json:"title" db:"title"`
	RawDate int64          `json:"date"`
	Date    time.Time      `db:"date"`
	Site    dat.NullString `json:"site" db:"site"`
}

type RawNomination struct {
	ID            int64          `db:"id"`
	MaleCount     int            `db:"male_count"`
	FemaleCount   int            `db:"female_count"`
	Type          string         `db:"type"`
	MinClass      dat.NullString `db:"min_class"`
	MaxClass      dat.NullString `db:"max_class"`
	MinJnjClass   dat.NullString `db:"min_jnj_class"`
	MaxJnjClass   dat.NullString `db:"max_jnj_class"`
	CompetitionID int64          `json:"competitionId" db:"competition_id"`
	Value         string         `json:"value" db:"value"`
}
