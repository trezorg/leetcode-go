package parentheses

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	last := 0
	var o1, o2, o3, c1, c2, c3 byte = '(', '{', '[', ')', '}', ']'
	stack := make([]byte, len(s), len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case o1, o2, o3:
			stack[last] = s[i]
			last++
		case c1:
			if last <= 0 || stack[last-1] != o1 {
				return false
			}
			last--
		case c2:
			if last <= 0 || stack[last-1] != o2 {
				return false
			}
			last--
		case c3:
			if last <= 0 || stack[last-1] != o3 {
				return false
			}
			last--
		}
	}
	return last == 0
}
