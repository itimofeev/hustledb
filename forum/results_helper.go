package forum

import (
	"github.com/itimofeev/hustlesa/util"
	"strings"
)

const judgeLetters = "ABCDEFGHI"
const classicClasses = "abcde"
const jnjClasses = "bgrsmsch"

// 1 (A) - Милованов Александр
func parseJudge(str string) *Judge {
	util.CheckMatchesRegexp("\\d \\(.\\) - [\\W ]+", str)

	letter := string(str[strings.Index(str, "(")+1])
	title := string(str[strings.Index(str, "- ")+2:])

	return &Judge{
		Letter: fixLetter(letter),
		Title:  strings.TrimSpace(title),
	}
}

func fixLetter(letter string) string {
	util.CheckOk(len(letter) == 1, "Not letter: "+letter)
	util.CheckOk(strings.Contains(judgeLetters, letter), "Unknown letter: "+letter)
	return letter
}

//1 место-№562-Беликов Александр Валерьевич(дебют,AlphaDance,E)-Егорова Юлия Викторовна(8463,AlphaDance,E)
//28-34 место-№504-Потапов Николай Олегович(7008,Движение,D,Bg)
func parsePlace(str string) *Place {
	//var dancerTitle
	//var dancerId
	//var clubs
	//var classClassic
	//var classJnj
	//
	//strings.Split(str, " место-")

	placeFrom, placeTo := parsePlaces(str[:strings.Index(str, "место")-1])

	sinceNumber := str[strings.Index(str, "№")+3:] //UTF staff
	numberString := sinceNumber[:strings.Index(sinceNumber, "-")]
	dancerString := sinceNumber[strings.Index(sinceNumber, "-")+1:]

	dancer1, dancer2 := parseDancers(dancerString)

	return &Place{
		PlaceFrom: placeFrom,
		PlaceTo:   placeTo,
		Number:    util.Atoi(numberString),
		Dancer1:   dancer1,
		Dancer2:   dancer2,
	}
}

func parsePlaces(str string) (int, int) {
	str = strings.TrimSpace(str)
	split := strings.Split(str, "-")

	if len(split) == 1 {
		return util.Atoi(str), util.Atoi(str)
	} else {
		return util.Atoi(split[0]), util.Atoi(split[1])
	}
}

func parseDancers(str string) (*Dancer, *Dancer) {
	split := strings.Split(str, "-")
	if len(split) == 1 {
		return parseDancer(str), nil
	} else {
		return parseDancer(split[0]), parseDancer(split[1])
	}
}

//Беликов Александр Валерьевич(дебют,AlphaDance,E)
//Потапов Николай Олегович(7008,Движение,Ivara,D,Bg)
func parseDancer(str string) *Dancer {
	title := str[:strings.Index(str, "(")]
	info := str[strings.Index(str, "(")+1:]

	dancerCodeStr := info[:strings.Index(info, ",")]
	var dancerId int
	if dancerCodeStr != "дебют" {
		dancerId = util.Atoi(dancerCodeStr)
	}

	sinceClubs := info[strings.Index(info, ",")+1 : len(info)-1]
	split := strings.Split(sinceClubs, ",")
	var clubs []string
	for _, clubOrClass := range split {
		if !isClass(clubOrClass) {
			clubs = append(clubs, clubOrClass)
		}
	}

	return &Dancer{
		Id:    dancerId,
		Title: title,
		Clubs: clubs,
	}
}

func isClass(letter string) bool {
	return strings.Contains(jnjClasses, strings.ToLower(letter)) || strings.Contains(classicClasses, strings.ToLower(letter))
}
