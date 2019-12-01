package wordbreak2

func _contains(s string, wordDict []string) bool {
	for _, str := range wordDict {
		if str == s {
			return true
		}
	}
	return false
}

func _wordBreak(s string, wordDict []string, idx int, cache *map[int][]string) []string {
	start := idx
	maxIdx := len(s) - 1
	current := make([]string, 0)
	for idx <= maxIdx {
		word := s[start : idx+1]
		if _contains(word, wordDict) {
			if idx == maxIdx {
				current = append(current, word)
				(*cache)[start] = current
				return current
			}
			values, ok := (*cache)[idx+1]
			if !ok {
				values = _wordBreak(s, wordDict, idx+1, cache)
			}
			for j := 0; j < len(values); j++ {
				current = append(current, word+" "+values[j])
			}
		}
		idx++
	}
	(*cache)[start] = current
	return current
}

func wordBreak(s string, wordDict []string) []string {
	cache := map[int][]string{}
	return _wordBreak(s, wordDict, 0, &cache)
}
