package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var data = []struct {
	s1     string
	s2     string
	result bool
}{
	{s1: "ab", s2: "eidbaooo", result: true},
	{s1: "ab", s2: "eidboaoo", result: false},
	{s1: "abc", s2: "bdddddddddbbbbbbbsdsbacbbbbbdsdbbbbbbbbbsdfsfsdfsdfsdddddfbbbbbb", result: true},
	{s1: "abc", s2: "bdddddddddcab", result: true},
	{s1: "abc", s2: "bdddddddddca", result: false},
	{s1: "ab", s2: "a", result: false},
	{s1: "ab", s2: "ba", result: true},
}

func TestCheckInclusion(t *testing.T) {
	for _, v := range data {
		t.Run(fmt.Sprintf("%s-%s", v.s1, v.s2), func(t *testing.T) {
			require.Equal(t, v.result, checkInclusion(v.s1, v.s2))
		})
	}

}

func BenchmarkCheckInclusion(b *testing.B) {

	for _, v := range data {
		b.Run(fmt.Sprintf("%s-%s", v.s1, v.s2), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				checkInclusion(v.s1, v.s2)
			}
		})
	}

}
