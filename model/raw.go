package model

import "gopkg.in/mgutz/dat.v1"

type RawParsingResults struct {
	Clubs   *[]RawClub
	Dancers *[]RawDancer
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
	ID          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	PrevSurname string `json:"prevSurname" db:"prev_surname"`
	PairClass   string `json:"pairClass" db:"pair_class"`
	JnjClass    string `json:"jnjClass" db:"jnj_class"`
	Sex         string `json:"sex" db:"sex"`
	Source      string `json:"source" db:"source"`
}
