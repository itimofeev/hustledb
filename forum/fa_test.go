package forum

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

var (
	beginType        = reflect.TypeOf(&BeginState{})
	judgeTeamType    = reflect.TypeOf(&JudgeTeamState{})
	placeType        = reflect.TypeOf(&PlacesState{})
	techType         = reflect.TypeOf(&TechnicalState{})
	techPreFinalType = reflect.TypeOf(&TechnicalPrepFinalState{})
	techFinalType    = reflect.TypeOf(&TechnicalFinalState{})
)

func TestTechnicalState_ProcessLine(t *testing.T) {
	fr := &ForumResults{}

	begin := &BeginState{}
	judge := &JudgeTeamState{}
	place := &PlacesState{}
	tech := &TechnicalState{}
	techPre := &TechnicalPrepFinalState{}
	techFinal := &TechnicalFinalState{}

	assert.Equal(t, beginType, reflect.TypeOf(begin.ProcessLine(fr, "21.00 Окончание турнира  ")))

	assert.Equal(t, judgeTeamType, reflect.TypeOf(begin.ProcessLine(fr, "Результаты турнира:")))
	assert.Equal(t, judgeTeamType, reflect.TypeOf(begin.ProcessLine(fr, "Результаты турнира, 1 отделение:")))
	assert.Equal(t, judgeTeamType, reflect.TypeOf(judge.ProcessLine(fr, "1 (A) - Милованов Александр")))
	assert.Equal(t, placeType, reflect.TypeOf(judge.ProcessLine(fr, "C класс. Участвовало пар: 6")))
	assert.Equal(t, placeType, reflect.TypeOf(place.ProcessLine(fr, "1/8 финала")))
	assert.Equal(t, placeType, reflect.TypeOf(place.ProcessLine(fr, "2 место-№740-Ларин Максим Геннадьевич(5944,Движение,C)-Тимофеева Юлия Андреевна(1956,Движение,C)")))

	assert.Equal(t, techType, reflect.TypeOf(place.ProcessLine(fr, "Технические результаты:")))

	assert.Equal(t, techType, reflect.TypeOf(tech.ProcessLine(fr, "")))
	assert.Equal(t, techType, reflect.TypeOf(tech.ProcessLine(fr, "DnD Beginner0 (ДЕВУШКИ): 1/4 финала")))
	assert.Equal(t, techType, reflect.TypeOf(tech.ProcessLine(fr, "  567   | AB|CDE ==> выход в 1/2 финала")))
	assert.Equal(t, techPreFinalType, reflect.TypeOf(tech.ProcessLine(fr, "C класс: ФИНАЛ")))
	assert.Equal(t, techPreFinalType, reflect.TypeOf(techPre.ProcessLine(fr, "--------+-------------------+--------+-------------------+--------+-------------------+--------+------+---------")))
	assert.Equal(t, techPreFinalType, reflect.TypeOf(techPre.ProcessLine(fr, "        | Места за 1-й.     │ Место  │ Места за 2-й.     │ Место  │ Места за 3-й.     │ Место  │Сумма │Итоговое ")))
	assert.Equal(t, techFinalType, reflect.TypeOf(techPre.ProcessLine(fr, "  637   | 2 2 1 2 5         │     2  │                   │        │                   │        │    2 │    2")))
	assert.Equal(t, techFinalType, reflect.TypeOf(techPre.ProcessLine(fr, "  007   ¦ 1 3 1 1 2         ¦     1  ¦                   ¦        ¦                   ¦        ¦   01 ¦    1")))

	assert.Equal(t, techFinalType, reflect.TypeOf(techFinal.ProcessLine(fr, "  599   | 6 5 6 7 7         │     6  │                   │        │                   │        │    6 │    6")))
	assert.Equal(t, judgeTeamType, reflect.TypeOf(techFinal.ProcessLine(fr, "--------+-------------------+--------+-------------------+--------+-------------------+--------+------+---------")))
}

func TestParseMainTitle(t *testing.T) {
	fr := &ForumResults{}

	parseMainTitle(fr, "(2016-09-17) Открытие сезона 2016-2017г., г. Москва. УТВЕРЖДЕНО РК АСХ")
	assert.Equal(t, "Открытие сезона 2016-2017г.", fr.Title)
	assert.Equal(t, "г. Москва. УТВЕРЖДЕНО РК АСХ", fr.Remaining)
	assert.Equal(t, time.Date(2016, 9, 17, 0, 0, 0, 0, time.UTC), fr.Date)

	parseMainTitle(fr, "    (2016-10-22) Осенний кубок клуба Движение, г. Москва, УТВЕРЖДЕНО РК АСХ  ")
	assert.Equal(t, "Осенний кубок клуба Движение", fr.Title)
	assert.Equal(t, "г. Москва. УТВЕРЖДЕНО РК АСХ", fr.Remaining)
	assert.Equal(t, time.Date(2016, 10, 22, 0, 0, 0, 0, time.UTC), fr.Date)

	parseMainTitle(fr, "(2016-11-05) Огни Владимира, г. Владимир, УТВЕРЖДЕНО РК АСХ")
	assert.Equal(t, "Огни Владимира", fr.Title)
	assert.Equal(t, "г. Владимир. УТВЕРЖДЕНО РК АСХ", fr.Remaining)
	assert.Equal(t, time.Date(2016, 11, 05, 0, 0, 0, 0, time.UTC), fr.Date)

	parseMainTitle(fr, "(2014-09-06) Открытие сезона (г.Москва), ДК Буревестник, м.Сокольники")
	assert.Equal(t, "Открытие сезона (г.Москва)", fr.Title)
	assert.Equal(t, "ДК Буревестник, м.Сокольники", fr.Remaining)
	assert.Equal(t, time.Date(2014, 9, 06, 0, 0, 0, 0, time.UTC), fr.Date)
}

func TestRegexpt(t *testing.T) {
	r := compileRegexp("^.*\\d+.*[│¦](.*\\d.*)+[│¦](.*[│¦])+.*\\d+$")

	line := "  007   ¦ 1 3 1 1 2         ¦     1  ¦                   ¦        ¦                   ¦        ¦   01 ¦    1"

	assert.True(t, r.MatchString(line))
}
