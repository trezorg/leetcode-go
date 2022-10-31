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

func checkPair(s []byte, p []byte, pair Pair) bool {

	ss := pair.s
	ps := pair.p

	if ss == len(s) && ps == len(p) {
		return true
	}
	if ps == len(p) && ss < len(s) {
		return false
	}
	if ss == len(s) {
		for _, ss := range p[ps:] {
			if ss == star {
				continue
			} else {
				return false
			}
		}
		return true
	}
	return false
}

func isMatchHelper(s []byte, p []byte) bool {

	pos := Pair{0, 0}
	res := checkPair(s, p, pos)
	if res {
		return res
	}

	sL := len(s)
	pL := len(p)

	stack := []Pair{pos}

	cache := map[Pair]bool{}

	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		pos := stack[lastIdx]
		stack = stack[:lastIdx]
		_, ok := cache[pos]

		if ok {
			continue
		}

		res := checkPair(s, p, pos)
		if res {
			return res
		}

		if pos.p >= pL || pos.s >= sL {
			continue
		}

		cache[pos] = true

		ss := s[pos.s]
		ps := p[pos.p]

		if ps != star {
			if ps == question || ps == ss {
				stack = append(stack, Pair{s: pos.s + 1, p: pos.p + 1})
			}
		} else if ps == star {
			stack = append(stack, Pair{s: pos.s, p: pos.p + 1})
			stack = append(stack, Pair{s: pos.s + 1, p: pos.p})
			stack = append(stack, Pair{s: pos.s + 1, p: pos.p + 1})
		}

	}

	return false
}

func isMatch(s string, p string) bool {
	ss := []byte(s)
	ps := splicePattern(p)
	return isMatchHelper(ss, ps)
}
