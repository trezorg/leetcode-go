package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

var benchmarkGroups [][]string

func canonicalGroups(groups [][]string) []string {
	result := make([]string, len(groups))
	for i, group := range groups {
		words := append([]string(nil), group...)
		sort.Strings(words)
		result[i] = strings.Join(words, ",")
	}
	sort.Strings(result)
	return result
}

func TestGroupAnagramsFast(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []string
		want  [][]string
	}{
		{"empty", nil, nil},
		{"single", []string{"a"}, [][]string{{"a"}}},
		{"mixed", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{"duplicates_and_empty", []string{"", "a", "", "a", "b"}, [][]string{{"", ""}, {"a", "a"}, {"b"}}},
		{"long_words", []string{strings.Repeat("a", 256), strings.Repeat("b", 256), strings.Repeat("a", 256)}, [][]string{{strings.Repeat("a", 256), strings.Repeat("a", 256)}, {strings.Repeat("b", 256)}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := groupAnagramsFast(append([]string(nil), tc.input...))
			if !reflect.DeepEqual(canonicalGroups(got), canonicalGroups(tc.want)) {
				t.Fatalf("groupAnagramsFast() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGroupAnagramsFastMatchesBaseline(t *testing.T) {
	for _, tc := range []struct {
		name          string
		groups        int
		wordsPerGroup int
	}{
		{"one_group", 1, 100},
		{"twenty_groups", 20, 50},
		{"unique_groups", 1000, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := benchmarkWords(tc.groups, tc.wordsPerGroup)
			want := canonicalGroups(groupAnagrams(append([]string(nil), input...)))
			got := canonicalGroups(groupAnagramsFast(append([]string(nil), input...)))
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("groupAnagramsFast() = %v, want %v", got, want)
			}
		})
	}
}

func benchmarkWords(groups, wordsPerGroup int) []string {
	words := make([]string, 0, groups*wordsPerGroup)
	for group := range groups {
		key := [4]byte{
			'a' + byte(group%5),
			'f' + byte(group/5%5),
			'k' + byte(group/25%5),
			'p' + byte(group/125%8),
		}
		for word := range wordsPerGroup {
			var rotated [4]byte
			for i := range rotated {
				rotated[i] = key[(i+word)%len(key)]
			}
			words = append(words, string(rotated[:]))
		}
	}
	return words
}

func benchmarkGroupAnagrams(b *testing.B, group func([]string) [][]string) {
	for _, tc := range []struct {
		name          string
		groups        int
		wordsPerGroup int
	}{
		{"one_group", 1, 1000},
		{"twenty_groups", 20, 50},
		{"unique_groups", 1000, 1},
	} {
		b.Run(tc.name, func(b *testing.B) {
			input := benchmarkWords(tc.groups, tc.wordsPerGroup)
			work := make([]string, len(input))
			b.ReportAllocs()
			for b.Loop() {
				b.StopTimer()
				copy(work, input)
				b.StartTimer()
				benchmarkGroups = group(work)
			}
		})
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	benchmarkGroupAnagrams(b, groupAnagrams)
}

func BenchmarkGroupAnagramsFast(b *testing.B) {
	benchmarkGroupAnagrams(b, groupAnagramsFast)
}
