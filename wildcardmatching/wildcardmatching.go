package wildcardmatching

const (
	star     = byte('*')
	question = byte('?')
)

// Pair struct
type Pair struct {
	p int
	s int
}

func splicePattern(p string) []byte {
	b := []byte(p)
	if len(p) < 2 {
		return []byte(p)
	}
	i, j := 0, 1
	for j < len(b) {
		fs := b[i]
		ss := b[j]
		if fs == star && ss == star {
			j++
		} else {
			if j-i > 1 {
				b[i+1] = ss
			}
			i++
			j++
		}
	}
	return b[:i+1]
}

func checkAndReturn(s []byte, p []byte, cache *map[Pair]bool) bool {
	pair := Pair{s: len(s), p: len(p)}
	val, ok := (*cache)[pair]
	if ok {
		return val
	}
	res := isMatchHelper(s, p, cache)
	(*cache)[pair] = res
	return res
}

func isMatchHelper(s []byte, p []byte, cache *map[Pair]bool) bool {

	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(p) == 0 && len(s) != 0 {
		return false
	}
	if len(s) == 0 {
		for _, ss := range p {
			if ss == star {
				continue
			} else {
				return false
			}
		}
		return true
	}

	ps := p[0]
	ss := s[0]

	if ps != star {
		if ps == question || ps == ss {
			return checkAndReturn(s[1:], p[1:], cache)
		}
		(*cache)[Pair{s: len(s[1:]), p: len(p[1:])}] = false
		return false
	} else if ps == star {
		return checkAndReturn(s, p[1:], cache) || checkAndReturn(s[1:], p[1:], cache) || checkAndReturn(s[1:], p, cache)
	}

	return false
}

func isMatch(s string, p string) bool {
	ss := []byte(s)
	ps := splicePattern(p)
	cache := map[Pair]bool{}
	return isMatchHelper(ss, ps, &cache)
}
