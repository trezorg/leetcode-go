package wildcardmatching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplicePattern1(t *testing.T) {
	s := "******"
	ss := splicePattern(s)
	assert.Len(t, ss, 1)
}

func TestSplicePattern2(t *testing.T) {
	s := "******b"
	ss := splicePattern(s)
	assert.Len(t, ss, 2)
}

func TestSplicePattern3(t *testing.T) {
	s := "**a*c**b"
	ss := splicePattern(s)
	assert.Len(t, ss, 6)
	assert.Equal(t, "*a*c*b", string(ss))
}

func TestSplicePattern4(t *testing.T) {
	s := "*?????"
	ss := splicePattern(s)
	assert.Len(t, ss, 6)
}

func TestSplicePattern5(t *testing.T) {
	s := "*?????**"
	ss := splicePattern(s)
	assert.Len(t, ss, 7)
	assert.Equal(t, "*?????*", string(ss))
}

func TestSplicePattern6(t *testing.T) {
	s := "abcdef"
	ss := splicePattern(s)
	assert.Len(t, ss, 6)
	assert.Equal(t, s, string(ss))
}

func TestSplicePattern7(t *testing.T) {
	s := "??????"
	ss := splicePattern(s)
	assert.Len(t, ss, 6)
	assert.Equal(t, "??????", string(ss))
}

func TestSplicePattern8(t *testing.T) {
	s := "b**bb**a**bba*b**a*bbb**aba***babbb*aa****aabb*bbb***a"
	ss := splicePattern(s)
	assert.Equal(t, "b*bb*a*bba*b*a*bbb*aba*babbb*aa*aabb*bbb*a", string(ss))
}

func TestMatching1(t *testing.T) {
	s := "aa"
	p := "a"
	assert.False(t, isMatch(s, p))
}

func TestMatching2(t *testing.T) {
	s := "aa"
	p := "*"
	assert.True(t, isMatch(s, p))
}

func TestMatching3(t *testing.T) {
	s := ""
	p := "*"
	assert.True(t, isMatch(s, p))
}

func TestMatching4(t *testing.T) {
	s := "cb"
	p := "?a"
	assert.False(t, isMatch(s, p))
}

func TestMatching5(t *testing.T) {
	s := "adceb"
	p := "*a*b"
	assert.True(t, isMatch(s, p))
}

func TestMatching6(t *testing.T) {
	s := "acdcb"
	p := "a*c?b"
	assert.False(t, isMatch(s, p))
}

func TestMatching7(t *testing.T) {
	s := "acdcb"
	p := "?????"
	assert.True(t, isMatch(s, p))
}

func TestMatching8(t *testing.T) {
	s := "acdcb"
	p := "??????"
	assert.False(t, isMatch(s, p))
}

func TestMatching9(t *testing.T) {
	s := "mississippi"
	p := "m??*ss*?i*pi"
	assert.False(t, isMatch(s, p))
}

func TestMatching10(t *testing.T) {
	s := "babbbbaabababaabbababaababaabbaabababbaaababbababaaaaaabbabaaaabababbabbababbbaaaababbbabbbbbbbbbbaabbb"
	p := "b**bb**a**bba*b**a*bbb**aba***babbb*aa****aabb*bbb***a"
	assert.False(t, isMatch(s, p))
}

func BenchmarkMatching(b *testing.B) {
	s := "babbbbaabababaabbababaababaabbaabababbaaababbababaaaaaabbabaaaabababbabbababbbaaaababbbabbbbbbbbbbaabbb"
	p := "b**bb**a**bba*b**a*bbb**aba***babbb*aa****aabb*bbb***a"
	for i := 0; i < b.N; i++ {
		assert.False(b, isMatch(s, p))
	}
}
