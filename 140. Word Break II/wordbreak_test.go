package wordbreak2

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
)

func TestWordBreak(t *testing.T) {

	s := "leetcode"
	wordDict := []string{"leet", "code"}
	assert.Equal(t, []string{"leet code"}, wordBreak(s, wordDict))

	s = "applepenapple"
	wordDict = []string{"apple", "pen"}
	assert.Equal(t, []string{"apple pen apple"}, wordBreak(s, wordDict))

	s = "applepenapple"
	wordDict = []string{"apple", "pen", "applepen"}
	assert.Equal(t, []string{"apple pen apple", "applepen apple"}, wordBreak(s, wordDict))

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	assert.Equal(t, []string{}, wordBreak(s, wordDict))

	s = "applepenword"
	wordDict = []string{"apple", "applepen", "word"}
	assert.Equal(t, []string{"applepen word"}, wordBreak(s, wordDict))

	s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	assert.Len(t, wordBreak(s, wordDict), 0)

	s = "catsanddog"
	wordDict = []string{"cat", "cats", "and", "sand", "dog"}
	assert.Equal(t, []string{"cat sand dog", "cats and dog"}, wordBreak(s, wordDict))

	s = "pineapplepenapple"
	wordDict = []string{"apple", "pen", "applepen", "pine", "pineapple"}
	assert.Equal(t, []string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"}, wordBreak(s, wordDict))

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	assert.Equal(t, []string{}, wordBreak(s, wordDict))

}

func BenchmarkWordBreak1(b *testing.B) {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	for i := 0; i <= b.N; i++ {
		assert.Len(b, wordBreak(s, wordDict), 0)
	}
}

func BenchmarkWordBreak2(b *testing.B) {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	for i := 0; i <= b.N; i++ {
		assert.Len(b, wordBreak(s, wordDict), 0)
	}
}

func BenchmarkWordBreak3(b *testing.B) {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	for i := 0; i <= b.N; i++ {
		assert.Len(b, wordBreak(s, wordDict), 0)
	}
}

func BenchmarkWordBreak4(b *testing.B) {
	s := strings.Repeat("pineapplepenapple", 10)
	wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	for i := 0; i <= b.N; i++ {
		wordBreak(s, wordDict)
	}
}
