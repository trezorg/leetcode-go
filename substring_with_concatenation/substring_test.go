package substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscring1(t *testing.T) {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	assert.Equal(t, []int{0, 9}, findSubstring(s, words))
}

func TestSubscring2(t *testing.T) {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word","good","best","word"}
	assert.Equal(t, []int{}, findSubstring(s, words))
}

func TestSubscring3(t *testing.T) {
	s := "ooooooooo"
	words := []string{"oo","oo"}
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, findSubstring(s, words))
}

func TestSubscring4(t *testing.T) {
	s := "foobarfbarfoofoobarfbarfoo"
	words := []string{"foo","bar", "oob", "arf"}
	assert.Equal(t, []int{1, 14}, findSubstring(s, words))
}
