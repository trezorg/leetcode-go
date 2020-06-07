package parentheses

import "sort"

func minRemoveToMakeValid(s string) string {
	buf := []byte(s)
	stack := make([]int, 0)
	remove := make([]int, 0)
	for idx, b := range buf {
		if b == '(' {
			stack = append(stack, idx)
		} else if b == ')' {
			if len(stack) == 0 {
				remove = append(remove, idx)
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	stack = append(stack, remove...)
	if len(stack) == 0 {
		return s
	}
	sort.Ints(stack)
	stackIdx := 0
	start := stack[stackIdx]
	for i := start + 1; i < len(buf); {
		if stackIdx+1 < len(stack) && stack[stackIdx+1] == i {
			i++
			stackIdx++
			continue
		}
		buf[start], buf[i] = buf[i], buf[start]
		start++
		i++
	}

	return string(buf[:start])

}
