package substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscring1(t *testing.T) {
	s := "ADOBECODEBANC"
	p := "ABC"
	assert.Equal(t, "BANC", minWindow(s, p))
}

func TestSubscring2(t *testing.T) {
	s := "ADOBECODEBANC"
	p := "ABCD"
	assert.Equal(t, "ADOBEC", minWindow(s, p))
}

func TestSubscring3(t *testing.T) {
	s := "XXXABYYCDOBECODEBANC"
	p := "ABC"
	assert.Equal(t, "BANC", minWindow(s, p))
}

func TestSubscring4(t *testing.T) {
	s := "YYYYYYYYYYYYYYYYABDDDDDDDDDDD"
	p := "ABC"
	assert.Equal(t, "", minWindow(s, p))
}

func TestSubscring5(t *testing.T) {
	s := ""
	p := ""
	assert.Equal(t, "", minWindow(s, p))
}

func TestSubscring6(t *testing.T) {
	s := "xx"
	p := ""
	assert.Equal(t, "", minWindow(s, p))
}

func TestSubscring7(t *testing.T) {
	s := "XYBAAAAAACABYC"
	p := "ABC"
	assert.Equal(t, "CAB", minWindow(s, p))
}

func TestSubscring8(t *testing.T) {
	s := "A"
	p := "AA"
	assert.Equal(t, "", minWindow(s, p))
}

func TestSubscring9(t *testing.T) {
	s := "AA"
	p := "AA"
	assert.Equal(t, "AA", minWindow(s, p))
}

func TestSubscring10(t *testing.T) {
	s := "AA"
	p := "AAA"
	assert.Equal(t, "", minWindow(s, p))
}

func TestSubscring11(t *testing.T) {
	s := "ABDCABC"
	p := "ABC"
	assert.Equal(t, "CAB", minWindow(s, p))
}
