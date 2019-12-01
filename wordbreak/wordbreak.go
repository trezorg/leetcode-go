package wordbreak

func _contains(s string, wordDict []string) bool {
	for _, str := range wordDict {
		if str == s {
			return true
		}
	}
	return false
}

func _wordBreak(s string, wordDict []string, idx int, cache *map[int]bool) bool {
	start := idx
	maxIdx := len(s) - 1
	for idx <= maxIdx {
		word := s[start : idx+1]
		if _contains(word, wordDict) {
			if idx == maxIdx {
				return true
			}
			_, ok := (*cache)[idx+1]
			if !ok {
				res := _wordBreak(s, wordDict, idx+1, cache)
				if res {
					return true
				}
			}
		}
		idx++
	}
	(*cache)[start] = false
	return false
}

func wordBreak(s string, wordDict []string) bool {
	cache := map[int]bool{}
	return _wordBreak(s, wordDict, 0, &cache)
}

