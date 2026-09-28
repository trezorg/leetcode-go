package main

import (
	"bytes"
	"slices"
)

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	var count [26]int
	for i := 0; i < len(a); i++ {
		count[a[i]-'a']++
		count[b[i]-'a']--
	}

	for _, n := range count {
		if n != 0 {
			return false
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	borders := []int{0}
	j := 0 // конец формируемой группы; там же место для следующего совпадения

	for j < len(strs) {
		word := strs[j]

		for i := j; i < len(strs); i++ {
			if isAnagram(word, strs[i]) {
				strs[j], strs[i] = strs[i], strs[j]
				j++
			}
		}

		borders = append(borders, j)
	}

	groups := make([][]string, 0, len(borders)-1)
	for k := 1; k < len(borders); k++ {
		groups = append(groups, strs[borders[k-1]:borders[k]])
	}
	return groups
}

func groupAnagramsFast(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	type signatureWord struct {
		word string
		key  [26]byte
	}
	pairs := make([]signatureWord, len(strs))
	for i, word := range strs {
		pairs[i].word = word
		for j := 0; j < len(word); j++ {
			pairs[i].key[word[j]-'a']++
		}
	}

	slices.SortFunc(pairs, func(a, b signatureWord) int {
		return bytes.Compare(a.key[:], b.key[:])
	})

	groups := make([][]string, 0)
	start := 0
	for i, pair := range pairs {
		strs[i] = pair.word
		if i > 0 && pair.key != pairs[i-1].key {
			groups = append(groups, strs[start:i])
			start = i
		}
	}
	return append(groups, strs[start:])
}
