package model

type DancerInfo struct {
	RawDancer

	Results []RawCompetitionResult `json:"results"`
	Clubs   []RawClub              `json:"clubs"`
}
