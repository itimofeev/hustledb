package parser

import (
	"bitbucket.org/Axxonsoft/axxoncloudgo/util"
	"encoding/json"
	"github.com/itimofeev/hustlesa/model"
	"io/ioutil"
)

func Parse(dirName string) *model.RawParsingResults {
	clubs := parseClubs(dirName + "clubs.json")
	dancers := parseDancers(dirName + "dancers.json")
	return &model.RawParsingResults{
		Clubs:   clubs,
		Dancers: dancers,
	}
}

func parseClubs(fileName string) *[]model.RawClub {
	data, err := ioutil.ReadFile(fileName)
	util.CheckErr(err, "Read file: "+fileName)

	clubs := make([]model.RawClub, 0)

	json.Unmarshal(data, &clubs)

	return &clubs
}

func parseDancers(fileName string) *[]model.RawDancer {
	data, err := ioutil.ReadFile(fileName)
	util.CheckErr(err, "Read file: "+fileName)

	dancers := make([]model.RawDancer, 0)

	json.Unmarshal(data, &dancers)

	return &dancers
}
