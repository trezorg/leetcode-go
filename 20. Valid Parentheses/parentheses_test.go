package parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParentheses(t *testing.T) {

	type value struct {
		data   string
		result bool
	}

	data := []value{
		{"[]", true},
		{"{}", true},
		{"{}]", false},
		{"[{}]", true},
		{"[]{}[]()", true},
		{"[{{([])}}]", true},
		{"[{{([])}{]", false},
		{"([)]", false},
		{"([{{}}])", true},
		{"([{)", false},
		{"", false},
	}

	for _, v := range data {
		t.Run(v.data, func(t *testing.T) {
			assert.Equal(t, isValid(v.data), v.result)
		})
	}

}
