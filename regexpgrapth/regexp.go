package regexp

import (
	"fmt"
)

const (
	dot            = rune('.')
	star           = rune('*')
	plus           = rune('+')
	parenthesLeft  = rune('(')
	parenthesRight = rune(')')
)

var (
	excludeSymbols = func() map[byte]struct{} {
		res := map[byte]struct{}{}
		res[star] = struct{}{}
		res[plus] = struct{}{}
		res[parenthesLeft] = struct{}{}
		res[parenthesRight] = struct{}{}
		return res
	}()
)

// Token interface
type Token interface {
	read(data string, pos int) int
	add(token ...Token) Token
	or(token ...Token) Token
	value() string
}

// AnySymbolToken for one of possible tokens
type AnySymbolToken struct {
	excludeSymbols map[byte]bool
	val            string
}

func (t AnySymbolToken) read(data string, pos int) int {
	if pos >= len(data) {
		return 0
	}
	_, ok := t.excludeSymbols[data[pos]]
	if ok {
		return 0
	}
	return 1
}

func (t AnySymbolToken) add(token ...Token) Token {
	tokens := []Token{t}
	return ConcatToken{tokens: append(tokens, token...)}
}

func (t AnySymbolToken) or(token ...Token) Token {
	tokens := []Token{t}
	return ConcatToken{tokens: append(tokens, token...)}
}

func (t AnySymbolToken) value() string {
	return t.val
}

// SymbolToken one token
type SymbolToken struct {
	symbol byte
}

func (t SymbolToken) read(data string, pos int) int {
	if pos >= len(data) {
		return 0
	}
	if t.symbol == data[pos] {
		return 1
	}
	return 0
}

// ConcatToken group of tokens
type ConcatToken struct {
	tokens []Token
}

func (t ConcatToken) read(data string, pos int) int {
	if pos >= len(data) {
		return 0
	}
	read := 0
	for _, token := range t.tokens {
		n := token.read(data, pos+read)
		if n == 0 {
			return read
		}
		read += n
	}
	return read
}

func (t ConcatToken) value() string {
	return ""
}

func (t ConcatToken) add(token ...Token) Token {
	return t
}

func (t ConcatToken) or(token ...Token) Token {
	return t
}

// SplitToken struct
type SplitToken struct {
	tokens []Token
}

func (t SplitToken) read(data string, pos int) int {
	if pos >= len(data) {
		return 0
	}
	for _, token := range t.tokens {
		n := token.read(data, pos)
		if n > 0 {
			return n
		}
	}
	return 0
}

func (t SplitToken) value() string {
	return ""
}

func (t SplitToken) add(token ...Token) Token {
	return t
}

func (t SplitToken) or(token ...Token) Token {
	return t
}

func parseTokens(data string) ([][]byte, error) {
	s := []byte(data)
	var res [][]byte
	for len(s) > 0 {
		n := 0
		for _, token := range allowedTokens {
			n = token.read(s)
			if n > 0 {
				res = append(res, s[:n])
				s = s[n:]
				break
			}
		}
		if n == 0 {
			return res, fmt.Errorf("cannot parse string")
		}
	}
	return res, nil
}
