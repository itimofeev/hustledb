package forum

import (
	"fmt"
	"testing"
)

func TestTechnicalState_ProcessLine(t *testing.T) {
	initRegexps()
	//fr := &ForumResults{}
	//state := &TechnicalPrepFinalState{
	//}

	line := "        | Места за 1-й.     │ Место  │ Места за 2-й.     │ Место  │ Места за 3-й.     │ Место  │Сумма │Итоговое "
	line2 := "  687   | 1 5 4 2 6         │     5  │ 4 4 3 2 5         │     4  │                   │        │    9 │    4"
	//newState := state.ProcessLine(fr, line)
	//fmt.Printf("!!!%T\n", newState)//TODO remove

	fmt.Printf("!!!%+v\n", techFinalResult.MatchString(line))  //TODO remove
	fmt.Printf("!!!%+v\n", techFinalResult.MatchString(line2)) //TODO remove
}
