package parentheses

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParentheses1(t *testing.T) {
	s := "lee(t(c)o)de)"
	res := "lee(t(c)o)de"
	assert.Equal(t, res, minRemoveToMakeValid(s))
}

func TestParentheses2(t *testing.T) {
	s := "))(("
	res := ""
	assert.Equal(t, res, minRemoveToMakeValid(s))
}

func TestParentheses3(t *testing.T) {
	s := "(a(b(c)d)"
	res := "a(b(c)d)"
	assert.Equal(t, res, minRemoveToMakeValid(s))
}

func TestParentheses4(t *testing.T) {
	s := ""
	res := ""
	assert.Equal(t, res, minRemoveToMakeValid(s))
}

func TestParentheses5(t *testing.T) {
	s := "((((((()))))))"
	res := "((((((()))))))"
	assert.Equal(t, res, minRemoveToMakeValid(s))
}

func TestParentheses6(t *testing.T) {
	s := "test((a()))oops((((((v))"
	res := "test((a()))oops((v))"
	assert.Equal(t, res, minRemoveToMakeValid(s))
}

func TestParentheses7(t *testing.T) {
	s := ")test(("
	res := "test"
	assert.Equal(t, res, minRemoveToMakeValid(s))
}

func TestParentheses8(t *testing.T) {
	s := strings.Repeat("(", 10000)
	res := ""
	assert.Equal(t, res, minRemoveToMakeValid(s))
}
