package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var data = []struct {
	s1     string
	s2     string
	result []int
}{
	{s1: "abc", s2: "cbaebabacd", result: []int{0, 6}},
	{s1: "ab", s2: "abab", result: []int{0, 1, 2}},
}

func TestCheckInclusion(t *testing.T) {
	for _, v := range data {
		t.Run(fmt.Sprintf("%s-%s", v.s2, v.s1), func(t *testing.T) {
			require.Equal(t, v.result, findAnagrams(v.s2, v.s1))
		})
	}

}

func BenchmarkCheckInclusion(b *testing.B) {

	for _, v := range data {
		b.Run(fmt.Sprintf("%s-%s", v.s2, v.s1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findAnagrams(v.s2, v.s1)
			}
		})
	}

}
