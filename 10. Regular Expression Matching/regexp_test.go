package regexp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseRegex(t *testing.T) {
	s := "a*c*..*b"
	regexp, err := parseRegexp(s)
	assert.NoError(t, err)
	require.Len(t, regexp.nodes, 5)
	firstNode := regexp.nodes[0]
	secondNode := regexp.nodes[1]
	thirdNode := regexp.nodes[2]
	fourthNode := regexp.nodes[3]
	fifthNode := regexp.nodes[4]
	assert.Equal(t, byte('a'), firstNode.symbol)
	assert.Equal(t, 0, firstNode.repeatMin)
	assert.Equal(t, infiniteRepeat, firstNode.repeatMax)
	assert.Equal(t, byte('c'), secondNode.symbol)
	assert.Equal(t, 0, secondNode.repeatMin)
	assert.Equal(t, infiniteRepeat, secondNode.repeatMax)

	assert.Equal(t, anySymbol, thirdNode.symbol)
	assert.Equal(t, 1, thirdNode.repeatMin)
	assert.Equal(t, 1, thirdNode.repeatMax)

	assert.Equal(t, anySymbol, fourthNode.symbol)
	assert.Equal(t, 0, fourthNode.repeatMin)
	assert.Equal(t, infiniteRepeat, fourthNode.repeatMax)

	assert.Equal(t, byte('b'), fifthNode.symbol)
	assert.Equal(t, 1, fifthNode.repeatMin)
	assert.Equal(t, 1, fifthNode.repeatMax)
}

func TestParseBlankRegex(t *testing.T) {
	s := ""
	regexp, err := parseRegexp(s)
	assert.NoError(t, err)
	assert.Len(t, regexp.nodes, 0)
}

func TestRegex(t *testing.T) {
	p := "a*c*..*b"
	s := "aaccxyyyyb"
	res := isMatch(s, p)
	assert.True(t, res)
	s = "aaccxyyyy"
	res = isMatch(s, p)
	assert.False(t, res)
}

func TestBlankRegex(t *testing.T) {
	p := ""
	s := ""
	res := isMatch(s, p)
	assert.True(t, res)
}

func TestRegexInput(t *testing.T) {
	s := "aa"
	p := "a"
	res := isMatch(s, p)
	assert.False(t, res)

	s = "aa"
	p = "a*"
	res = isMatch(s, p)
	assert.True(t, res)

	s = "ab"
	p = ".*"
	res = isMatch(s, p)
	assert.True(t, res)

	s = "aab"
	p = "c*a*b"
	res = isMatch(s, p)
	assert.True(t, res)

	s = "aab"
	p = "c*a*b"
	res = isMatch(s, p)
	assert.True(t, res)

	s = "mississippi"
	p = "mis*is*p*."
	res = isMatch(s, p)
	assert.False(t, res)

}

func TestRegexAnySymbol1(t *testing.T) {
	p := ".*.*.*"
	s := "aaccxyyyyb"
	res := isMatch(s, p)
	assert.True(t, res)
	s = "y"
	res = isMatch(s, p)
	assert.True(t, res)
	s = "yy"
	res = isMatch(s, p)
	assert.True(t, res)
	s = "yaya"
	res = isMatch(s, p)
	assert.True(t, res)
}

func TestRegexAnySymbol2(t *testing.T) {
	p := ".*.*c.*"
	s := "aaccxyyyyb"
	res := isMatch(s, p)
	assert.True(t, res)
	s = "y"
	res = isMatch(s, p)
	assert.False(t, res)
	s = "yay"
	res = isMatch(s, p)
	assert.False(t, res)
}

func TestRegexAnySymbol3(t *testing.T) {

	p := ".*.*.*"
	s := "a"
	res := isMatch(s, p)
	assert.True(t, res)

	p = "a*"
	s = ""
	res = isMatch(s, p)
	assert.True(t, res)

	p = "a"
	s = ""
	res = isMatch(s, p)
	assert.False(t, res)

	p = "a*"
	s = "a"
	res = isMatch(s, p)
	assert.True(t, res)

	p = "a*"
	s = "c"
	res = isMatch(s, p)
	assert.False(t, res)

	p = ".*"
	s = "c"
	res = isMatch(s, p)
	assert.True(t, res)

	p = ".*.*.*"
	s = ""
	res = isMatch(s, p)
	assert.True(t, res)

	p = "a*a*a*"
	s = ""
	res = isMatch(s, p)
	assert.True(t, res)

	p = ".*.*c.*"
	regexp, err := parseRegexp(p)
	assert.NoError(t, err)
	assert.Equal(t, 3, regexp.len())
	s = "c"
	res = isMatch(s, p)
	assert.True(t, res)

	p = ".*.*c.*a*c*"
	s = "c"
	res = isMatch(s, p)
	assert.True(t, res)

	p = "c.*a*c*t"
	s = "cat"
	res = isMatch(s, p)
	assert.True(t, res)

	p = "c.*a*c*"
	s = "cax"
	res = isMatch(s, p)
	assert.True(t, res)

	p = ".*a*c*"
	s = "ax"
	res = isMatch(s, p)
	assert.True(t, res)

	p = ".*a*c*x"
	s = "ax"
	res = isMatch(s, p)
	assert.True(t, res)

	p = ".*a*c*v"
	s = "ax"
	res = isMatch(s, p)
	assert.False(t, res)

	p = ".*a*c*."
	s = "ax"
	res = isMatch(s, p)
	assert.True(t, res)

	p = "c*"
	s = "ccc"
	res = isMatch(s, p)
	assert.True(t, res)

	p = "c*c*c*"
	s = "cccx"
	res = isMatch(s, p)
	assert.False(t, res)

}

func TestRegexAnySymbol4(t *testing.T) {

	p := ".*.."
	s := "a"

	regexp, err := parseRegexp(p)
	assert.NoError(t, err)
	assert.Equal(t, 3, regexp.len())
	dotToken := regexp.nodes[2]
	assert.Equal(t, 1, dotToken.repeatMin)
	assert.Equal(t, 1, dotToken.repeatMax)
	assert.Equal(t, anySymbol, dotToken.symbol)

	res := isMatch(s, p)
	assert.False(t, res)

}

func TestRegexAnySymbol5(t *testing.T) {
	s := "aaaaaaaaaaaaab"
	p := "a*a*a*a*a*a*a*a*a*a*c"
	res := isMatch(s, p)
	assert.False(t, res)
}

func TestRegexAnySymbol6(t *testing.T) {
	s := "aaaaaaaaaaaaab"
	p := "a*b*c*d*e*f*j*i*k*j*m"
	res := isMatch(s, p)
	assert.False(t, res)
}

func TestRegexAnySymbol7(t *testing.T) {
	s := "aaaaaaaaaaaaab"
	p := ".*.*.d*.*.*.v*.*.*.p*.*.*c"
	res := isMatch(s, p)
	assert.False(t, res)
}

func TestRegexAnySymbol8(t *testing.T) {
	s := "a"
	p := ".*..a*"
	res := isMatch(s, p)
	assert.False(t, res)
}

func TestParseRegex2(t *testing.T) {
	p := "a*a*a*a*a*a*a*a*a*a*c"
	regexp, err := parseRegexp(p)
	assert.NoError(t, err)
	assert.Equal(t, 2, regexp.len())
	symbolToken := regexp.nodes[1]
	assert.Equal(t, byte('c'), symbolToken.symbol)
	aToken := regexp.nodes[0]
	assert.Equal(t, byte('a'), aToken.symbol)
	assert.Equal(t, infiniteRepeat, aToken.repeatMax)
	assert.Equal(t, 0, aToken.repeatMin)

}

func BenchmarkTestRegexAnySymbol(b *testing.B) {
	s := "aaaaaaaaaaaaab"
	p := "a*a*a*a*a*a*a*a*a*a*c"
	for i := 0; i <= b.N; i++ {
		res := isMatch(s, p)
		assert.False(b, res)
	}
}
