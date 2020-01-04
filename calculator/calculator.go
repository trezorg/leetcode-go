package calculator

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type tokenType int

const (
	digits         = "0123456789"
	operatorSum    = "+"
	operatorSub    = "-"
	operatorMul    = "*"
	operatorDevide = "/"
	operatorMod    = "%"
	operators      = operatorSum + operatorSub + operatorMul + operatorDevide + operatorMod
	spaceSymbol    = rune(' ')
	groupStart     = rune('(')
	groupEnd       = rune(')')
	eof            = rune(0)

	tokenTypeInt tokenType = iota + 1
	tokenTypeOperator
	tokenTypeGroupStart
	tokenTypeGroupEnd
	tokenTypeError
)

var (
	operatorFuncs = map[string]calc{
		operatorSum: func(node1 node, node2 node) int {
			return node1.value() + node2.value()
		},
		operatorSub: func(node1 node, node2 node) int {
			return node1.value() - node2.value()
		},
		operatorMul: func(node1 node, node2 node) int {
			return node1.value() * node2.value()
		},
		operatorDevide: func(node1 node, node2 node) int {
			return node1.value() / node2.value()
		},
		operatorMod: func(node1 node, node2 node) int {
			return node1.value() % node2.value()
		},
	}
)

type stack struct {
	nodes []node
}

func (s *stack) push(node node) {
	(*s).items = append(s.nodes, node)
}

func (s *stack) pop() node {
	if len(s.nodes) == 0 {
		return nil
	}
	lastIdx := len(s.nodes) - 1
	node := s.nodes[lastIdx]
	(*s).items = s.nodes[:lastIdx]
	return node
}

func (s *stack) length() int {
	return len(s.nodes)
}

func (s *stack) get() node {
	if len(s.nodes) == 0 {
		return nil
	}
	lastIdx := len(s.nodes) - 1
	nodes := s.nodes[lastIdx]
	return item
}

type token struct {
	tp    tokenType
	value string
}

type lexer struct {
	input   string
	start   int
	pos     int
	width   int
	results chan token
}

func (l *lexer) run() <-chan token {
	go func() {
		for state := valueLexer(l); state != nil; {
			state = state(l)
		}
		close(l.results)
	}()
	return l.results
}

func (l *lexer) errorF(message string, args ...interface{}) stateFN {
	l.results <- token{tp: tokenTypeError, value: fmt.Sprintf(message, args...)}
	return func(*lexer) stateFN {
		return nil
	}
}

func newLexer(input string) *lexer {
	return &lexer{
		input:   input,
		results: make(chan token),
	}
}

func (l *lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
}

func (l *lexer) position() string {
	return fmt.Sprintf("%s - %s. Position: %d. Symbol: %c", l.input[:l.pos], l.input[l.pos:], l.pos, l.peek())
}

func (l *lexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) forward() {
	l.start = l.pos
	l.width = 0
}

func (l *lexer) read() string {
	res := l.input[l.start:l.pos]
	l.forward()
	return res
}

func (l *lexer) emit(tp tokenType) {
	l.results <- token{tp: tp, value: l.read()}
}

func (l *lexer) swallow(rune rune) {
	for l.next() == rune {
		l.forward()
	}
	l.backup()
}

func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptRune(rune rune) bool {
	if l.next() == rune {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) swallowOne(rune rune) {
	if l.next() == rune {
		l.forward()
	} else {
		l.backup()
	}
}

func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *lexer) acceptTill(rune rune) {
	for l.next() != rune {
	}
	l.backup()
}

type node interface {
	value() int
}

type calc func(node1 node, node2 node) int
type stateFN func(*lexer) stateFN

type leafNone struct {
	item int
}

func (l leafNone) String() string {
	return fmt.Sprintf("leafNone{item=%d}", l.item)
}

func (l leafNone) value() int {
	return l.item
}

type operatorNode struct {
	left     node
	right    node
	operator string
}

func (o operatorNode) String() string {
	return fmt.Sprintf("operatorNode{operator=%s, left: %s, right: %s}", o.operator, o.left, o.right)
}

func (o operatorNode) value() int {
	return operatorFuncs[o.operator](o.left, o.right)
}

func valueLexer(l *lexer) stateFN {
	l.swallow(spaceSymbol)
	if l.accept(digits) {
		l.acceptRun(digits)
		l.emit(tokenTypeInt)
		return operatorLexer
	}
	switch l.next() {
	case groupStart:
		l.emit(tokenTypeGroupStart)
		return valueLexer
	case eof:
		return nil
	default:
		return l.errorF("Parsing error: %s. Awaiting value or group", l.position())
	}
}

func operatorLexer(l *lexer) stateFN {
	l.swallow(spaceSymbol)
	if l.accept(operators) {
		l.emit(tokenTypeOperator)
		return valueLexer
	}
	switch k := l.next(); k {
	case groupEnd:
		l.emit(tokenTypeGroupEnd)
		return operatorLexer
	case eof:
		return nil
	default:
		fmt.Println(string(k))
		return l.errorF("Parsing error: %s. Awaiting operator", l.position())
	}
}

func toAst(tokens chan token) (node, error) {
	stack := stack{}
	for token := range tokens {
		switch token.tp {
		case tokenTypeInt:
			value, _ := strconv.Atoi(token.value)
			stack.push(leafNone{item: value})
		case tokenTypeOperator:
			switch token.value {
			case operatorSum, operatorSub:
				prev := stack.pop()
				stack.push(operatorNode{
					left:     prev,
					operator: token.value,
					right:    nil,
				})
			case operatorMul, operatorDevide, operatorMod:
				prev := stack.pop()
				stack.push(operatorNode{
					left:     prev,
					operator: token.value,
					right:    nil,
				})
			}
		}
	}
	return nil, nil
}

func parse(expression string) error {
	lexer := newLexer(expression)
	chn := lexer.run()
	for item := range chn {
		if item.tp == tokenTypeError {
			return fmt.Errorf(item.value)
		}
	}
	return nil
}
