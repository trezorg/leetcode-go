package substring

type wordDict map[string]int

func newWordDict(words []string) wordDict {
	m := make(wordDict)
	for _, w := range words {
		m[w]++
	}
	return m
}

func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(words) == 0 || len(s) == 0 {
		return res
	}
	lw := len(words[0])
	lws := len(words)
	lwss := lw * lws
	ls := len(s)
	lastI := ls - lwss + 1
	lastR := ls - lw + 1
	wd := newWordDict(words)
	for i := 0; i < lastI; {
		if wd[s[i:i+lw]] == 0 {
			i++
			continue
		}
		found := 1
		wd[s[i:i+lw]]--
		r := i + lw
		for r < lastR && wd[s[r:r+lw]] != 0 {
			wd[s[r:r+lw]]--
			found++
			r += lw
			if found == lws {
				break
			}
		}
		if found == lws {
			res = append(res, i)
		}
		i++
		wd = newWordDict(words)
	}
	return res
}
