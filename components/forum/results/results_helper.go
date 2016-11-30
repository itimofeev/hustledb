package forum

import (
	"fmt"
	"github.com/itimofeev/hustledb/components/util"
	"regexp"
	"strings"
)

const judgeLetters = "ABCDEFGHIJKL"
const classicClasses = "abcde"
const jnjClasses = "bgrsmsch"

// 1 (A) - Милованов Александр
func parseJudge(str string) *Judge {
	util.CheckMatchesRegexp("\\d \\(.\\) - [\\W ]+", str)

	letter := str[strings.Index(str, "(")+1 : strings.Index(str, ")")]
	title := str[strings.Index(str, "- ")+2:]

	return &Judge{
		Letter: fixJudgeLetter(letter),
		Title:  strings.TrimSpace(title),
	}
}

func fixJudgeLetter(letter string) string {
	letter = fixLetter(letter)
	util.CheckOk(len(letter) == 1, fmt.Sprintf("Not letter: '%s'", letter))
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
	str = strings.TrimLeft(str, "\"")
	split := strings.Split(str, "-")

	if len(split) == 1 {
		return util.Atoi(str), util.Atoi(str)
	} else {
		return util.Atoi(split[0]), util.Atoi(split[1])
	}
}

func parseDancers(str string) (*Dancer, *Dancer) {
	split := strings.Split(str, ")-")
	if len(split) == 1 {
		return parseDancer(str), nil
	} else {
		return parseDancer(split[0]), parseDancer(split[1])
	}
}

//Беликов Александр Валерьевич(дебют,AlphaDance,E)
//Потапов Николай Олегович(7008,Движение,Ivara,D,Bg)
func parseDancer(str string) *Dancer {
	str = strings.TrimLeft(str, "\"")
	str = strings.TrimRight(str, "\"")
	r, _ := regexp.Compile("\\(((\\d+)||(дебют)),")

	indexOpen := r.FindStringIndex(str)

	title := str[:indexOpen[0]]
	info := str[indexOpen[0]+1:]

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
		Id:    int64(dancerId),
		Title: title,
		Clubs: clubs,
	}
}

func fixLetter(letter string) string {
	switch letter {
	case "А":
		return "A"
	case "В":
		return "B"
	case "С":
		return "C"
	case "Д":
		return "D"
	case "Е":
		return "E"
	default:
		return letter
	}
}

func isClass(letter string) bool {
	letter = strings.TrimSpace(letter)
	letter = strings.ToLower(letter)
	return strings.Contains(jnjClasses, letter) || strings.Contains(classicClasses, letter)
}

func parseNominationTitleCount(line string) (string, int) {
	split := strings.Split(line, ".")
	if len(split) != 2 {
		return line, -1
	}

	countString := split[1]
	countString = strings.Replace(countString, "Участников:", "", 1)
	countString = strings.Replace(countString, "Участвовало", "", 1)
	countString = strings.Replace(countString, "пар:", "", 1)
	countString = strings.TrimSpace(countString)

	count := util.Atoi(countString)
	return split[0], count
}
