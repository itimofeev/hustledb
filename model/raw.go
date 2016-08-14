package model

import "gopkg.in/mgutz/dat.v1"

type RawParsingResults struct {
	Clubs       []RawClub
	Dancers     []RawDancer
	DancerClubs []RawDancerClub
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
	Surname    dat.NullString `db:"surname"`
	Patronymic dat.NullString `db:"patronymic"`
}

type RawDancerClub struct {
	DancerId  int64  `json:"dancerId" db:"dancer_id"`
	ClubNames string `json:"clubId"`
	ClubId    int64  `db:"club_id"`
}
