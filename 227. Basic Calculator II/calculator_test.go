package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse1(t *testing.T) {
	s := "1 + 2 + 3"
	assert.Equal(t, 6, calculate(s))
}

func TestParse2(t *testing.T) {
	s := "1 + 2 * 3"
	assert.Equal(t, 7, calculate(s))
}

func TestParse3(t *testing.T) {
	s := "1 + 2 * 3 / 6"
	assert.Equal(t, 2, calculate(s))
}

func TestParse4(t *testing.T) {
	s := "1 + 2 * 3 / 6 - 2"
	assert.Equal(t, 0, calculate(s))
}

func TestParse5(t *testing.T) {
	s := "1 + 2 * 3 / 6 - 2 / 2 - 1"
	assert.Equal(t, 0, calculate(s))
}

func TestParse6(t *testing.T) {
	s := " 3 +5/2"
	assert.Equal(t, 5, calculate(s))
}

func TestParse7(t *testing.T) {
	s := "0"
	assert.Equal(t, 0, calculate(s))
}

func TestParse8(t *testing.T) {
	s := " "
	assert.Equal(t, 0, calculate(s))
}

func TestParse9(t *testing.T) {
	s := " 0 "
	assert.Equal(t, 0, calculate(s))
}

func TestParse10(t *testing.T) {
	s := " 1 "
	assert.Equal(t, 1, calculate(s))
}

func TestParse11(t *testing.T) {
	s := "0"
	assert.Equal(t, 0, calculate(s))
}

func TestParse12(t *testing.T) {
	s := "14/3*2"
	assert.Equal(t, 8, calculate(s))
}

func TestParse13(t *testing.T) {
	s := "14/3+2"
	assert.Equal(t, 6, calculate(s))
}
