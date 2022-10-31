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
)
const (
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
	tokens []token
}

func (s *stack) push(token token) {
	(*s).tokens = append(s.tokens, token)
}

func (s *stack) pop() *token {
	if len(s.tokens) == 0 {
		return nil
	}
	lastIdx := len(s.tokens) - 1
	token := s.tokens[lastIdx]
	(*s).tokens = s.tokens[:lastIdx]
	return &token
}

func (s *stack) length() int {
	return len(s.tokens)
}

func (s *stack) get() *token {
	if len(s.tokens) == 0 {
		return nil
	}
	lastIdx := len(s.tokens) - 1
	token := s.tokens[lastIdx]
	return &token
}

type token struct {
	tp    tokenType
	value string
}

func (t *token) String() string {
	return fmt.Sprintf("token{type: %d, value: %s}", t.tp, t.value)
}

type lexer struct {
	input   string
	start   int
	pos     int
	width   int
	results []token
}

func (l *lexer) run() []token {
	for state := valueLexer(l); state != nil; {
		state = state(l)
	}
	return l.results
}

func (l *lexer) errorF(message string, args ...interface{}) stateFN {
	l.results = append(l.results, token{tp: tokenTypeError, value: fmt.Sprintf(message, args...)})
	return func(*lexer) stateFN {
		return nil
	}
}

func newLexer(input string) *lexer {
	return &lexer{
		input:   input,
		results: []token{},
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
	l.results = append(l.results, token{tp: tp, value: l.read()})
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
	addWithOperator(operator string, node node) (node, error)
	isGroup() bool
	setAsGroup()
}

type calc func(node1 node, node2 node) int
type stateFN func(*lexer) stateFN

type leafNode struct {
	item int
}

func (l *leafNode) String() string {
	return fmt.Sprintf("leafNode{item=%d}", l.item)
}

func (l *leafNode) value() int {
	return l.item
}

func (l *leafNode) isGroup() bool {
	return false
}

func (l *leafNode) setAsGroup() {
}

func (l *leafNode) addWithOperator(operator string, node node) (node, error) {
	switch operator {
	case operatorSum, operatorSub:
		return &operatorNode{operator: operator, left: l, right: node}, nil
	case operatorMul, operatorDevide, operatorMod:
		return &operatorNode{operator: operator, left: l, right: node, group: true}, nil
	}
	return nil, fmt.Errorf("unsupported operator: %s", operator)
}

type operatorNode struct {
	left     node
	right    node
	operator string
	group    bool
}

func (o *operatorNode) String() string {
	return fmt.Sprintf("operatorNode{operator=%s, left: %s, right: %s}", o.operator, o.left, o.right)
}

func (o *operatorNode) value() int {
	return operatorFuncs[o.operator](o.left, o.right)
}

func (o *operatorNode) isGroup() bool {
	return o.group
}

func (o *operatorNode) setAsGroup() {
	o.group = true
}

func (o *operatorNode) addWithOperator(operator string, node node) (node, error) {
	switch operator {
	case operatorSum, operatorSub:
		return &operatorNode{operator: operator, left: o, right: node}, nil
	case operatorMul, operatorDevide, operatorMod:
		if !o.isGroup() {
			o.right = &operatorNode{operator: operator, left: o.right, right: node}
			return o, nil
		}
		return &operatorNode{operator: operator, left: o, right: node, group: true}, nil
	}
	return nil, fmt.Errorf("unsupported operator: %s", operator)
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
	switch l.next() {
	case groupEnd:
		l.emit(tokenTypeGroupEnd)
		return operatorLexer
	case eof:
		return nil
	default:
		return l.errorF("Parsing error: %s. Awaiting operator", l.position())
	}
}

func toAst(tokens []token, start int) (node, int, error) {
	var s stack
	var opNode node
	var err error
	var i = start
	for i < len(tokens) {
		t := tokens[i]
		switch t.tp {
		case tokenTypeError:
			return nil, 0, fmt.Errorf(t.value)
		case tokenTypeInt:
			value, _ := strconv.Atoi(t.value)
			node := leafNode{item: value}
			if s.length() == 0 {
				opNode = &node
			} else {
				operatorItem := s.pop()
				opNode, err = opNode.addWithOperator(operatorItem.value, &node)
				if err != nil {
					return nil, 0, err
				}
				if s.length() > 0 {
					return nil, 0, fmt.Errorf("wrong tokens processing. stack: %v. current op: %s", s.tokens, opNode)
				}
			}
		case tokenTypeOperator:
			s.push(t)
		case tokenTypeGroupEnd:
			if s.length() > 0 {
				return nil, 0, fmt.Errorf("wrong tokens positions. stack: %v. current op: %s", s.tokens, opNode)
			}
			opNode.setAsGroup()
			return opNode, i, nil
		case tokenTypeGroupStart:
			var node node
			node, i, err = toAst(tokens, i+1)
			if err != nil {
				return nil, i, err
			}
			if s.length() == 0 {
				opNode = node
			} else {
				operatorItem := s.pop()
				opNode, err = opNode.addWithOperator(operatorItem.value, node)
				if err != nil {
					return nil, i, err
				}
				if s.length() > 0 {
					return nil, i, fmt.Errorf("wrong tokens processing. stack: %v. current op: %s", s.tokens, opNode)
				}
			}
		default:
			return nil, i, fmt.Errorf("unsupported operator: %s", t.value)
		}
		i++
	}
	if s.length() > 0 {
		return nil, i, fmt.Errorf("wrong tokens processing. stack: %v current op: %s", s.tokens, opNode)
	}
	return opNode, i, nil
}

func parse(expression string) (node, int, error) {
	lexer := newLexer(expression)
	channel := lexer.run()
	return toAst(channel, 0)
}

func calculate(s string) int {
	node, _, _ := parse(s)
	return node.value()
}
