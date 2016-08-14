package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, "", "")

	m := make(map[string]int64, 0)

	m["Красный уголок"] = 777

	assert.Equal(t, 777, m["Красный уголок"])
}
