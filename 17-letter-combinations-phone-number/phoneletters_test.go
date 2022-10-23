package phoneletters

import (
	"fmt"
	"testing"
	"sort"

	"github.com/stretchr/testify/assert"
)

func TestPhoneLetters(t *testing.T) {

	data := []struct {
		data   string
		result []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"2", []string{"a", "b", "c"}},
		{"234", []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%s", v.data), func(t *testing.T) {
			sort.Strings(v.result)
			res := letterCombinations(v.data)
			sort.Strings(res)
			assert.Equal(t, v.result, res)
		})
	}

}
