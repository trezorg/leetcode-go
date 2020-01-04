package calculator

import (
	"strconv"
	"strings"
)

const (
	plus   = byte('+')
	minus  = byte('-')
	mult   = byte('*')
	devide = byte('/')
	space  = byte(' ')
	digits = "0123456789"
)

type stack struct {
	tokens []int
}

func (s *stack) push(token int) {
	s.tokens = append(s.tokens, token)
}

func (s *stack) pop() int {
	if len(s.tokens) == 0 {
		return -1
	}
	lastIdx := len(s.tokens) - 1
	token := s.tokens[lastIdx]
	s.tokens = s.tokens[:lastIdx]
	return token
}

func (s *stack) length() int {
	return len(s.tokens)
}

func newStack(size int) *stack {
	return &stack{tokens: make([]int, 0, size)}
}

func op(second int, one int, sign byte) int {
	switch sign {
	case plus:
		return one + second
	case minus:
		return one - second
	case mult:
		return one * second
	case devide:
		return one / second
	default:
		panic("wrong sign: " + string(sign))
	}
}

func readDigits(s string, start int) (int, int) {
	var i int
	for i = start; i < len(s) && s[i] == space; i++ {
	}
	start = i
	for i = start; i < len(s) && strings.IndexByte(digits, s[i]) >= 0; i++ {
	}
	val := s[start:i]
	n, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return n, i
}

func calculate(s string) int {
	var sign byte
	st := newStack(2)
	for i := 0; i < len(s); {
		v := s[i]
		if v == space {
			i++
			continue
		}
		if v == plus || v == minus {
			if st.length() < 2 {
				sign = v
			} else {
				st.push(op(st.pop(), st.pop(), sign))
				sign = v
			}
			i++
		} else if v == mult || v == devide {
			if st.length() < 2 {
				sign = v
				i++
			} else {
				if sign == mult || sign == devide {
					st.push(op(st.pop(), st.pop(), sign))
					sign = v
					i++
				} else {
					var n int
					n, i = readDigits(s, i+1)
					st.push(op(n, st.pop(), v))
				}
			}
		} else {
			var n int
			n, i = readDigits(s, i)
			st.push(n)
		}
	}
	if st.length() == 1 {
		return st.pop()
	} else if st.length() == 0 {
		return 0
	}
	return op(st.pop(), st.pop(), sign)
}
