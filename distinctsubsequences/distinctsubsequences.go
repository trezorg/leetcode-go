package distinctsubsequences

type pos struct {
	s, t int
}


func _numDistinct(s string, t string, sIdx int, tIdx int, cache *map[pos]int) int {
	maxSidx := len(s) - 1
	maxTidx := len(t) - 1
	start := sIdx
	lt := len(t) - tIdx
	tm := t[tIdx]
	res := 0
	for sIdx + lt <= len(s) {
		sm := s[sIdx]
		if sm == tm {
			if lt == 1 {
				res++
			}
			if tIdx < maxTidx && sIdx < maxSidx {
				c := pos{sIdx + 1, tIdx + 1}
				rs, ok := (*cache)[c]
				if !ok {
					res += _numDistinct(s, t, sIdx + 1, tIdx + 1, cache)
				} else {
					res += rs
				}
			}
		}
		sIdx++
	}
	(*cache)[pos{start, tIdx}] = res
	return res
}

func numDistinct(s string, t string) int {
	cache := map[pos]int{}
	return _numDistinct(s, t, 0, 0, &cache)
}


func _numDistinct2(s string, t string) int {

	maxSidx := len(s) - 1
	maxTidx := len(t) - 1
	res := 0
	stack := []pos{}
	stack = append(stack, pos{0, 0})
	cache := map[pos]int{}

	for len(stack) > 0 {

		n := len(stack) - 1
		posV := stack[n]
		stack = stack[:n]

		sIdx, tIdx := posV.s, posV.t
		lt := len(t) - tIdx
		tm := t[tIdx]
	
		for sIdx + lt <= len(s) {
			sm := s[sIdx]
			if sm == tm {
				if lt == 1 {
					res++
				}
				if tIdx < maxTidx && sIdx < maxSidx {
					c := pos{sIdx + 1, tIdx + 1}
					rs, ok := cache[posV]
					if !ok {
						stack = append(stack, c)
					} else {
						res += rs
					}
				}
			}
			sIdx++
		}
		cache[posV] = res

	}
	return res
}

func numDistinct2(s string, t string) int {
	return _numDistinct2(s, t)
}