package parser

import (
	"bitbucket.org/Axxonsoft/axxoncloudgo/util"
	"encoding/json"
	"fmt"
	"github.com/itimofeev/hustlesa/model"
	"gopkg.in/mgutz/dat.v1"
	"io/ioutil"
	"log"
	"strings"
)

type parser struct {
	dirName     string
	clubs       []model.RawClub
	dancers     []model.RawDancer
	dancerClubs []model.RawDancerClub

	name2club map[string]int64
}

func (p *parser) Parse() model.RawParsingResults {
	p.name2club = make(map[string]int64)
	p.clubs = fixClubs(parseClubs(p.dirName + "clubs.json"))

	p.dancers = fixDancers(parseDancers(p.dirName + "dancers.json"))

	fillName2Club(p.name2club, p.clubs)

	dancerClubs := parseDancerClubs(p.dirName + "dancerClubs.json")
	p.dancerClubs = fixDancerClubs(dancerClubs, p.name2club)

	return model.RawParsingResults{
		Clubs:       p.clubs,
		Dancers:     p.dancers,
		DancerClubs: p.dancerClubs,
	}
}

func parseDancerName(name string) (string, string, *string) {
	split := strings.Split(name, " ")
	if !(len(split) == 2 || len(split) == 3) {
		log.Panic("Bad name " + name)
	}

	if len(split) == 2 {
		return split[0], split[1], nil
	}
	return split[0], split[1], &split[2]
}

func fixDancers(dancers []model.RawDancer) []model.RawDancer {
	for i, dancer := range dancers {
		dancers[i].Code = fmt.Sprintf("%05d", dancer.ID)
		surname, name, patronymic := parseDancerName(dancer.Title)
		dancers[i].Name = name
		dancers[i].Surname = surname
		dancers[i].Title = ""

		if dancers[i].Sex == "м" {
			dancers[i].Sex = "m"
		} else if dancers[i].Sex == "ж" {
			dancers[i].Sex = "f"
		} else {
			CheckErr(errors.New("bad sex "+dancers[i].Sex), "")
		}

		if patronymic != nil {
			dancers[i].Patronymic = dat.NullStringFrom(*patronymic)
		}
	}

	return dancers
}

func fixClubs(clubs []model.RawClub) []model.RawClub {
	maxClubId := findMaxClubId(clubs)
	clubs = append(clubs, model.RawClub{ID: 0, Name: "самост."})
	clubs = append(clubs, model.RawClub{ID: maxClubId + 1, Name: "Magnit"})
	clubs = append(clubs, model.RawClub{ID: maxClubId + 2, Name: "Intensity (г.Иваново)"})
	clubs = append(clubs, model.RawClub{ID: maxClubId + 3, Name: "Мартэ"})

	return clubs
}

func findMaxClubId(clubs []model.RawClub) int64 {
	var maxId int64 = clubs[0].ID
	for _, club := range clubs {
		if club.ID > maxId {
			maxId = club.ID
		}
	}
	return maxId
}

func fixDancerClubs(original []model.RawDancerClub, name2club map[string]int64) []model.RawDancerClub {
	dancerClubs := make([]model.RawDancerClub, 0, len(original)*2)
	for _, dc := range original {
		names := strings.Split(dc.ClubNames, ",")

		generated := generateDancerClubs(names, name2club, dc)

		dancerClubs = append(dancerClubs, generated...)
	}

	return dancerClubs
}
func generateDancerClubs(names []string, name2club map[string]int64, original model.RawDancerClub) []model.RawDancerClub {
	if len(names) == 1 {
		clubId, ok := name2club[strings.ToLower(names[0])]
		if !ok {
			log.Panic("Not found club name " + names[0])
		}
		original.ClubId = clubId
		return []model.RawDancerClub{original}
	}

	dancerClubs := make([]model.RawDancerClub, 0)
	for _, name := range names {
		club, ok := name2club[strings.ToLower(name)]
		if !ok {
			log.Panicf("Not found club name '%s', %+v", name, original)
		}

		dancerClub := model.RawDancerClub{ClubId: club, DancerId: original.DancerId, ClubNames: name}
		dancerClubs = append(dancerClubs, dancerClub)
	}

	return dancerClubs
}

func Parse(dirName string) model.RawParsingResults {
	p := parser{dirName: dirName}

	return p.Parse()
}
func fillName2Club(name2club map[string]int64, clubs []model.RawClub) {
	for _, club := range clubs {
		name2club[strings.ToLower(club.Name)] = club.ID
	}
}

func parseClubs(fileName string) []model.RawClub {
	data, err := ioutil.ReadFile(fileName)
	util.CheckErr(err, "Read file: "+fileName)

	clubs := make([]model.RawClub, 0)

	json.Unmarshal(data, &clubs)

	return clubs
}

func parseDancers(fileName string) []model.RawDancer {
	data, err := ioutil.ReadFile(fileName)
	util.CheckErr(err, "Read file: "+fileName)

	dancers := make([]model.RawDancer, 0)

	json.Unmarshal(data, &dancers)

	return dancers
}

func parseDancerClubs(fileName string) []model.RawDancerClub {
	data, err := ioutil.ReadFile(fileName)
	util.CheckErr(err, "Read file: "+fileName)

	dancerClubs := make([]model.RawDancerClub, 0)

	json.Unmarshal(data, &dancerClubs)

	return dancerClubs
}
