package distinctsubsequences


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistinctSubsequences1(t *testing.T) {

	s1 := "bbb"
	s2 := "bbb"
	assert.Equal(t, 1, numDistinct(s1, s2))

	s1 = "bbb"
	s2 = "bb"
	assert.Equal(t, 3, numDistinct(s1, s2))

	s1 = "bbbab"
	s2 = "bb"
	assert.Equal(t, 6, numDistinct(s1, s2))

	s1 = "bbbbab"
	s2 = "bb"
	assert.Equal(t, 10, numDistinct(s1, s2))

	s1 = "bbbbab"
	s2 = "bba"
	assert.Equal(t, 6, numDistinct(s1, s2))

	s1 = "rabbbit"
	s2 = "rabbit"
	assert.Equal(t, 3, numDistinct(s1, s2))

	s1 = "babgbag"
	s2 = "bag"
	assert.Equal(t, 5, numDistinct(s1, s2))

	s1 = "bbbbb"
	s2 = "bbb"
	assert.Equal(t, 10, numDistinct(s1, s2))

	s1 = "bbbab"
	s2 = "bbb"
	assert.Equal(t, 4, numDistinct(s1, s2))

}

func TestDistinctSubsequences2(t *testing.T) {

	s1 := "bbb"
	s2 := "bbb"
	assert.Equal(t, 1, numDistinct2(s1, s2))

	s1 = "bbb"
	s2 = "bb"
	assert.Equal(t, 3, numDistinct2(s1, s2))

	s1 = "bbbab"
	s2 = "bb"
	assert.Equal(t, 6, numDistinct2(s1, s2))

	s1 = "bbbbab"
	s2 = "bb"
	assert.Equal(t, 10, numDistinct2(s1, s2))

	s1 = "bbbbab"
	s2 = "bba"
	assert.Equal(t, 6, numDistinct2(s1, s2))

	s1 = "rabbbit"
	s2 = "rabbit"
	assert.Equal(t, 3, numDistinct2(s1, s2))

	s1 = "babgbag"
	s2 = "bag"
	assert.Equal(t, 5, numDistinct2(s1, s2))

}

func BenchmarkDistinctSubsequences1(b *testing.B) {
	s1 := "adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc"
	s2 := "bcddceeeebecbc"	
	for i := 0; i <= b.N; i++ {
		numDistinct(s1, s2)
	}
}

func enchmarkDistinctSubsequences2(b *testing.B) {
	s1 := "adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc"
	s2 := "bcddceeeebecbc"	
	for i := 0; i <= b.N; i++ {
		numDistinct2(s1, s2)
	}
}
