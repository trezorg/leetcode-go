package longestpalindromicsubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {

	data := []struct {
		data   string
		result string
	}{
		{"babad", "aba"},
		{"cddb", "dd"},
		{"cddb", "dd"},
		{"ddb", "dd"},
		{"abcddcba", "abcddcba"},
		{"arcddcba", "cddc"},
		{"aaaacdcdfdfsffqwfqfewqfefwef", "aaaa"},
		{"aacdcdfdfsffqwfqfewqfefwef", "fqf"},
		{"abcabcabc", "b"},
		{"xyzazzzzzzzdddd", "zzzzzzz"},
		{"bb", "bb"},
		{"bbb", "bbb"},
		{"cccc", "cccc"},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%s-%s", v.data, v.result), func(t *testing.T) {
			assert.Equal(t, v.result, longestPalindrome(v.data))
		})
	}

}
