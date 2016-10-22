package forum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseJudge(t *testing.T) {
	judge := parseJudge("1 (A) - Милованов Александр ")

	assert.Equal(t, "A", judge.Letter)
	assert.Equal(t, "Милованов Александр", judge.Title)

	judge = parseJudge("2 (В) - Дубровин Игорь")

	assert.Equal(t, "B", judge.Letter)
	assert.Equal(t, "Дубровин Игорь", judge.Title)
}

//1 место-№562-Беликов Александр Валерьевич(дебют,AlphaDance,E)-Егорова Юлия Викторовна(8463,AlphaDance,E)
//28-34 место-№504-Потапов Николай Олегович(7008,Движение,D,Bg)
func TestParsePlace(t *testing.T) {
	place := parsePlace("1 место-№562-Беликов Александр Валерьевич(дебют,AlphaDance,E)-Егорова Юлия Викторовна(8463,AlphaDance,E)")
	assert.Equal(t, 1, place.PlaceFrom)
	assert.Equal(t, 1, place.PlaceTo)
	assert.Equal(t, 562, place.Number)
	assert.Equal(t, "Беликов Александр Валерьевич", place.Dancer1.Title)
	assert.Equal(t, 0, place.Dancer1.Id)
	assert.Equal(t, []string{"AlphaDance"}, place.Dancer1.Clubs)
	assert.Equal(t, "Егорова Юлия Викторовна", place.Dancer2.Title)
	assert.Equal(t, 8463, place.Dancer2.Id)
	assert.Equal(t, []string{"AlphaDance"}, place.Dancer2.Clubs)

	place = parsePlace("28-34 место-№504-Потапов Николай Олегович(7008,Движение,Ivara,D,Bg)")
	assert.Equal(t, 28, place.PlaceFrom)
	assert.Equal(t, 34, place.PlaceTo)
	assert.Equal(t, 504, place.Number)
	assert.Equal(t, "Потапов Николай Олегович", place.Dancer1.Title)
	assert.Equal(t, 7008, place.Dancer1.Id)
	assert.Equal(t, []string{"Движение", "Ivara"}, place.Dancer1.Clubs)

	place = parsePlace("19-23 место-№556-Козлов Корней Викторович(8432,Хастл-Центр,E)-Воропаева Марина Валериевна(дебют,Хастл-Центр,E)")
	assert.Equal(t, "Козлов Корней Викторович", place.Dancer1.Title)
	assert.Equal(t, []string{"Хастл-Центр"}, place.Dancer1.Clubs)

	place = parsePlace("72-85 место-№669-Тугаринова (Рико) Наталья Александровна(8514,Движение,E,Bg)")
	assert.Equal(t, "Тугаринова (Рико) Наталья Александровна", place.Dancer1.Title)
	assert.Equal(t, []string{"Движение"}, place.Dancer1.Clubs)

	place = parsePlace("19-23 место-№573-Мосолов Валерий Валентинович(8655,Дизайн (г.Курск),E)-Савельева Ирина Юрьевна(7331,Дизайн (г.Курск),E)")
	assert.Equal(t, []string{"Дизайн (г.Курск)"}, place.Dancer1.Clubs)

}
