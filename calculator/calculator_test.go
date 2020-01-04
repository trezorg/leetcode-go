package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse1(t *testing.T) {
	s := "1 + 2 + 3"
	assert.NoError(t, parse(s))
}

func TestParse2(t *testing.T) {
	s := "(1 + 2) + 3"
	assert.NoError(t, parse(s))
}

func TestParse3(t *testing.T) {
	s := "(1 + 2) + (3"
	assert.NoError(t, parse(s))
}

func TestParse4(t *testing.T) {
	s := "(1 + 2)   +      (3 + 2)"
	assert.NoError(t, parse(s))
}
