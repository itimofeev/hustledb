package parser

import (
	"bitbucket.org/Axxonsoft/axxoncloudgo/util"
	"encoding/json"
	"github.com/itimofeev/hustlesa/model"
	"io/ioutil"
)

func ParseClubs(fileName string) *[]model.RawClub {
	data, err := ioutil.ReadFile(fileName)
	util.CheckErr(err, "Read file: "+fileName)

	clubs := make([]model.RawClub, 0)

	json.Unmarshal(data, &clubs)

	return &clubs
}
