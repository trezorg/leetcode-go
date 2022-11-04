package main

func checkValidStringBruteForce(s string) bool {
	if len(s) == 0 {
		return false
	}

	type T struct {
		i int
		f int
	}

	ss_open, ss_closed, wildcard := byte('('), byte(')'), byte('*')
	length := 0
	for i := 0; i < len(s); i++ {
		if s[i] == wildcard {
			length++
		}
	}
	stack := make([]T, 0, length*2+1)
	stack = append(stack, T{i: 0, f: 0})

	for len(stack) > 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	l:
		for i := t.i; i < len(s); i++ {
			switch s[i] {
			case wildcard:
				stack = append(stack, T{i: i + 1, f: t.f + 1}) // (
				if t.f-1 >= 0 {
					stack = append(stack, T{i: i + 1, f: t.f - 1}) // )
				}
				continue
			case ss_open:
				t.f++
			case ss_closed:
				t.f--
			}
			if t.f < 0 {
				break l
			}
		}
		if t.f == 0 {
			return true
		}
	}

	return false
}

func checkValidStringBruteForceRecursive(s string) bool {
	if len(s) == 0 {
		return false
	}

	ss_open, ss_closed, wildcard := byte('('), byte(')'), byte('*')

	var dfs func(idx, f int) bool

	dfs = func(idx, f int) bool {
		for i := idx; i < len(s); i++ {
			switch s[i] {
			case wildcard:
				if dfs(i+1, f) || dfs(i+1, f+1) {
					return true
				}
				if f-1 >= 0 && dfs(i+1, f-1) {
					return true
				}
				continue
			case ss_open:
				f++
			case ss_closed:
				f--
			}
			if f < 0 {
				return false
			}
		}
		return f == 0
	}

	return dfs(0, 0)
}

func checkValidString(s string) bool {
	if len(s) == 0 {
		return false
	}

	ss_open, ss_closed, wildcard := byte('('), byte(')'), byte('*')
	lBalance, rBalance := 0, 0

	for i, j := 0, len(s)-1; i < len(s) && j >= 0; i, j = i+1, j-1 {
		switch s[i] {
		case wildcard, ss_open:
			lBalance++
		case ss_closed:
			lBalance--
		}
		switch s[j] {
		case wildcard, ss_closed:
			rBalance++
		case ss_open:
			rBalance--
		}
		if lBalance < 0 {
			return false
		}
		if rBalance < 0 {
			return false
		}
	}
	return true

}
