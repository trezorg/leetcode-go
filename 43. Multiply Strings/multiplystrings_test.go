package multiplystrings

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplyOne1(t *testing.T) {
	s := "1234"
	c := 4
	ss, _ := strconv.Atoi(s)
	res := strconv.Itoa(ss * c)
	assert.Equal(t, res, multiplyOne(s, c))
}

func TestMultiplyOne2(t *testing.T) {
	s := "1111"
	c := 1
	ss, _ := strconv.Atoi(s)
	res := strconv.Itoa(ss * c)
	assert.Equal(t, res, multiplyOne(s, c))
}

func TestMultiplyOne3(t *testing.T) {
	s := "9999"
	c := 9
	ss, _ := strconv.Atoi(s)
	res := strconv.Itoa(ss * c)
	assert.Equal(t, res, multiplyOne(s, c))
}

func TestAddStrings1(t *testing.T) {
	s1 := "9999"
	s2 := "9999"
	ss1, _ := strconv.Atoi(s1)
	ss2, _ := strconv.Atoi(s2)
	res := strconv.Itoa(ss1 + ss2)
	assert.Equal(t, res, addStrings(s1, s2))
}

func TestAddStrings2(t *testing.T) {
	s1 := "9999"
	s2 := "1111"
	ss1, _ := strconv.Atoi(s1)
	ss2, _ := strconv.Atoi(s2)
	res := strconv.Itoa(ss1 + ss2)
	assert.Equal(t, res, addStrings(s1, s2))
}

func TestMultiplyStrings1(t *testing.T) {
	s1 := "9999"
	s2 := "0"
	ss1, _ := strconv.Atoi(s1)
	ss2, _ := strconv.Atoi(s2)
	res := strconv.Itoa(ss1 * ss2)
	assert.Equal(t, res, multiply(s1, s2))
}

func TestMultiplyStrings2(t *testing.T) {
	s1 := "9999"
	s2 := "9999"
	ss1, _ := strconv.Atoi(s1)
	ss2, _ := strconv.Atoi(s2)
	res := strconv.Itoa(ss1 * ss2)
	assert.Equal(t, res, multiply(s1, s2))
}

func BenchmarkMultiplyStrings(b *testing.B) {
	s1 := strings.Repeat("123456789", 110)
	s2 := strings.Repeat("123456789", 110)
	for i := 0; i <= b.N; i++ {
		multiply(s1, s2)
	}
}
