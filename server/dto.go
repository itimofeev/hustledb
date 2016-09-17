package server

import "time"

type DancerProfileDto struct {
	ID    int64  `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Code  string `json:"code" db:"code"`

	ClassicClass string `json:"classicClass" db:"classic_class"`
	JnjClass     string `json:"jnjClass" db:"jnj_class"`

	Clubs []ClubDto `json:"clubs"`

	Results []ResultDto `json:"results"`
}

type ClubDto struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type ResultDto struct {
	ID int64 `json:"id" db:"id"`

	CompetitionID    int64     `json:"competitionId" db:"comp_id"`
	CompetitionTitle string    `json:"competitionTitle" db:"comp_title"`
	CompetitionDate  time.Time `json:"competitionDate" db:"comp_date"`

	NominationID    int64  `json:"nominationId" db:"nom_id"`
	NominationTitle string `json:"nominationTitle" db:"nom_title"`

	ResultString string `json:"resultString" db:"result"`

	IsJNJ bool `json:"isJnj"db:"is_jnj"`

	Points int    `json:"points" db:"points"`
	Class  string `json:"class" db:"class"`

	Place string `json:"place" db:"place"`
}
