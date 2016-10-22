package forum

import (
	"github.com/itimofeev/hustlesa/util"
	"strings"
)

const judgeLetters = "ABCDEFGHI"

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
	//var placeFrom int
	//var placeTo int
	//var number
	//var dancerTitle
	//var dancerId
	//var clubs
	//var classClassic
	//var classJnj
	//
	//strings.Split(str, " место-")

	return nil
}
