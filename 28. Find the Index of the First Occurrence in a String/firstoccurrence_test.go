package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var data = []struct {
	s1     string
	s2     string
	result int
}{
	{s1: "sadbutsad", s2: "sad", result: 0},
	{s1: "ssadbutsad", s2: "sad", result: 1},
	{s1: "leetcode", s2: "leeto", result: -1},
	{s1: "mississippi", s2: "pi", result: 9},
	{s1: "mississippi", s2: "issi", result: 1},
	{s1: "mississippi", s2: "assi", result: -1},
	{s1: "mississippi", s2: "iss", result: 1},
	{s1: "baabbaaaaaaabbaaaaabbabbababaabbabbbbbabbabbbbbbabababaabbbbbaaabbbbabaababababbbaabbbbaaabbaababbbaabaabbabbaaaabababaaabbabbababbabbaaabbbbabbbbabbabbaabbbaa", s2: "bbaaaababa", result: 107},
	{s1: "baab", s2: "cca", result: -1},
	{s1: "abc", s2: "c", result: 2},
	{s1: "", s2: "c", result: -1},
}

func TestStrStr(t *testing.T) {
	for _, v := range data {
		t.Run(fmt.Sprintf("%s-%s", v.s1, v.s2), func(t *testing.T) {
			require.Equal(t, v.result, strStr(v.s1, v.s2))
		})
	}

}
