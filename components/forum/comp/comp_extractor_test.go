package comp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTryFindCIty(t *testing.T) {
	assert.Equal(t, "Москва", *tryFindCity("г.Москва, Утверждено РК АСХ"))
	assert.Equal(t, "Москва", *tryFindCity("г.  Москва, Утверждено РК АСХ"))
	assert.Equal(t, "Минск", *tryFindCity("Беларусь, г. Минск, УТВЕРЖДЕНО РК АСХ"))
	assert.Equal(t, "Москваа", *tryFindCity("г.Москваа"))
	assert.Equal(t, "Москваа", *tryFindCity("helo ther, г.Москваа"))
	assert.Equal(t, "Москваа", *tryFindCity("helo ther, г.Москваа, 123 РК АСХ , ыроыва"))
	assert.Equal(t, "Саратов", *tryFindCity("УТВЕРЖДЕНО, г. Саратов"))
	assert.Equal(t, "Москва", *tryFindCity("г.Москва. УТВЕРЖДЕНО РК АСХ"))

	assert.Nil(t, tryFindCity("г., Утверждено РК АСХ"))
}
