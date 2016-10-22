package forum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseJudge(t *testing.T) {
	judge := parseJudge("1 (A) - Милованов Александр ")

	assert.Equal(t, "A", judge.Letter)
	assert.Equal(t, "Милованов Александр", judge.Title)
}
