package forum

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
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
	initRegexps()
	fr := &ForumResults{}

	begin := &BeginState{}
	judge := &JudgeTeamState{}
	place := &PlacesState{}
	tech := &TechnicalState{}
	techPre := &TechnicalPrepFinalState{}
	techFinal := &TechnicalFinalState{}

	assert.Equal(t, beginType, reflect.TypeOf(begin.ProcessLine(fr, "21.00 Окончание турнира  ")))

	assert.Equal(t, judgeTeamType, reflect.TypeOf(begin.ProcessLine(fr, "Результаты турнира:")))
	assert.Equal(t, placeType, reflect.TypeOf(judge.ProcessLine(fr, "C класс. Участвовало пар: 6")))
	assert.Equal(t, placeType, reflect.TypeOf(place.ProcessLine(fr, "2 место-№740-Ларин Максим Геннадьевич(5944,Движение,C)-Тимофеева Юлия Андреевна(1956,Движение,C)")))

	assert.Equal(t, techType, reflect.TypeOf(place.ProcessLine(fr, "Технические результаты:")))

	assert.Equal(t, techType, reflect.TypeOf(tech.ProcessLine(fr, "")))
	assert.Equal(t, techType, reflect.TypeOf(tech.ProcessLine(fr, "DnD Beginner0 (ДЕВУШКИ): 1/4 финала")))
	assert.Equal(t, techType, reflect.TypeOf(tech.ProcessLine(fr, "  567   | AB|CDE ==> выход в 1/2 финала")))
	assert.Equal(t, techPreFinalType, reflect.TypeOf(tech.ProcessLine(fr, "C класс: ФИНАЛ")))
	assert.Equal(t, techPreFinalType, reflect.TypeOf(techPre.ProcessLine(fr, "--------+-------------------+--------+-------------------+--------+-------------------+--------+------+---------")))
	assert.Equal(t, techPreFinalType, reflect.TypeOf(techPre.ProcessLine(fr, "        | Места за 1-й.     │ Место  │ Места за 2-й.     │ Место  │ Места за 3-й.     │ Место  │Сумма │Итоговое ")))
	assert.Equal(t, techFinalType, reflect.TypeOf(techPre.ProcessLine(fr, "  637   | 2 2 1 2 5         │     2  │                   │        │                   │        │    2 │    2")))

	assert.Equal(t, techFinalType, reflect.TypeOf(techFinal.ProcessLine(fr, "  599   | 6 5 6 7 7         │     6  │                   │        │                   │        │    6 │    6")))
	assert.Equal(t, judgeTeamType, reflect.TypeOf(techFinal.ProcessLine(fr, "--------+-------------------+--------+-------------------+--------+-------------------+--------+------+---------")))
}
