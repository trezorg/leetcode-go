package parentheness

import (
	"unsafe"
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := make([]byte, n*2)
	results := make([]string, 0)

	addResult := func() {
		result := make([]byte, n*2)
		copy(result, res)
		s := *(*string)(unsafe.Pointer(&result))
		results = append(results, s)
	}

	var dfs func(cur, open, closed int)

	dfs = func(cur, open, closed int) {
		if closed == 0 && open == 0 {
			addResult()
			return
		}
		if closed == open {
			res[cur] = '('
			dfs(cur+1, open-1, closed)
		} else {
			if open > 0 {
				res[cur] = '('
				dfs(cur+1, open-1, closed)
			}
			if closed > 0 {
				res[cur] = ')'
				dfs(cur+1, open, closed-1)
			}
		}
	}

	dfs(0, n, n)

	return results

}

func generateParenthesisUpdated(n int) []string {
	if n == 0 {
		return []string{}
	}
	results := make([]string, 0)

	type pair struct {
		closed int
		open   int
		res    string
	}

	stack := []pair{{closed: n, open: n - 1, res: "("}}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if p.closed == 0 && p.open == 0 {
			results = append(results, p.res)
		} else {
			if p.open < p.closed {
				stack = append(stack, pair{open: p.open, closed: p.closed - 1, res: p.res + ")"})
			}
			if p.open != 0 {
				stack = append(stack, pair{open: p.open - 1, closed: p.closed, res: p.res + "("})
			}
		}
	}
	return results

}
