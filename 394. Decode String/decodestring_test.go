package decodestring


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeString(t *testing.T) {
	s := "3[a]2[bc]"
	assert.Equal(t, "aaabcbc", decodeString(s))
	s = "3[a2[c]]"
	assert.Equal(t, "accaccacc", decodeString(s))
	s = "2[abc]3[cd]ef"
	assert.Equal(t, "abcabccdcdcdef", decodeString(s))
	s = "[abc]3[cd]ef"
	assert.Equal(t, "abccdcdcdef", decodeString(s))
	s = "[abc][cd][ef]"
	assert.Equal(t, "abccdef", decodeString(s))
	s = "a"
	assert.Equal(t, "a", decodeString(s))
	s = ""
	assert.Equal(t, "", decodeString(s))
	s = "3[a]2[b4[F]c]"
	assert.Equal(t, "aaabFFFFcbFFFFc", decodeString(s))
}
