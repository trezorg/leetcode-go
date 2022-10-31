package booleanexpression

import (
	"fmt"
	"unicode/utf8"
)

const (
	trueSymbol  = rune('t')
	falseSymbol = rune('f')
	commaSymbol = rune(',')
	orSymbol    = rune('|')
	notSymbol   = rune('!')
	andSymbol   = rune('&')
	groupStart  = rune('(')
	groupEnd    = rune(')')
	eof         = rune(0)
)

type stack struct {
	items []item
}

func (s *stack) push(item item) {
	(*s).items = append(s.items, item)
}

func (s *stack) pop() item {
	if len(s.items) == 0 {
		return nil
	}
	lastIdx := len(s.items) - 1
	item := s.items[lastIdx]
	(*s).items = s.items[:lastIdx]
	return item
}

func (s *stack) length() int {
	return len(s.items)
}

func (s *stack) get() item {
	if len(s.items) == 0 {
		return nil
	}
	lastIdx := len(s.items) - 1
	item := s.items[lastIdx]
	return item
}

type lexer struct {
	input string
	start int
	pos   int
	width int
	items stack
}

func newLexer(input string) *lexer {
	return &lexer{
		input: input,
	}
}

func (l *lexer) addItem(item item) {
	l.items.push(item)
}

func (l *lexer) closeItem() {
	lastItem := l.items.pop()
	if lastItem != nil {
		prevItem := l.items.get()
		if prevItem != nil {
			prevItem.add(lastItem)
		} else {
			l.items.push(lastItem)
		}
	}
}

func (l *lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		return eof
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
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

type stateFN func(*lexer) stateFN

func outerLex(l *lexer) stateFN {
	switch s := l.next(); s {
	case notSymbol:
		l.forward()
		return notState
	case orSymbol:
		l.forward()
		return orState
	case andSymbol:
		l.forward()
		return andState
	case commaSymbol:
		l.forward()
		return outerLex
	case groupEnd:
		l.forward()
		l.closeItem()
		return outerLex
	case eof:
		l.closeItem()
		return eofState
	default:
		return l.errorF("Outer lexer. bad syntax: %s - %s. Position: %d", l.input[:l.pos], l.input[l.pos:], l.pos)
	}
}

func innerLex(l *lexer) stateFN {
	switch l.next() {
	case trueSymbol:
		l.forward()
		l.addItem(&boolItem{item: true})
		l.closeItem()
		return innerLex
	case falseSymbol:
		l.forward()
		l.addItem(&boolItem{item: false})
		l.closeItem()
		return innerLex
	case commaSymbol:
		l.forward()
		return innerLex
	case groupEnd:
		l.forward()
		l.closeItem()
		return outerLex
	case notSymbol, orSymbol, andSymbol:
		l.backup()
		return outerLex
	default:
		return l.errorF("Inner lexer. bad syntax: %s - %s. Position: %d", l.input[:l.pos], l.input[l.pos:], l.pos)
	}
}

func eofState(l *lexer) stateFN {
	return nil
}

func (l *lexer) errorF(format string, args ...interface{}) stateFN {
	message := fmt.Sprintf(format, args...)
	return func(*lexer) stateFN {
		panic(message)
	}
}

func notState(l *lexer) stateFN {
	switch l.next() {
	case groupStart:
		l.addItem(&notBoolItem{item: nil})
		return innerLex
	default:
		return l.errorF("bad syntax: %s - %s. Position: %d", l.input[:l.pos], l.input[l.pos:], l.pos)
	}
}

func orState(l *lexer) stateFN {
	switch l.next() {
	case groupStart:
		l.forward()
		l.addItem(&orGroupBoolItem{})
		return innerLex
	default:
		return l.errorF("bad syntax: %s - %s. Position: %d", l.input[:l.pos], l.input[l.pos:], l.pos)
	}
}

func andState(l *lexer) stateFN {
	switch l.next() {
	case groupStart:
		l.forward()
		l.addItem(&andGroupBoolItem{})
		return innerLex
	default:
		return l.errorF("bad syntax: %s - %s. Position: %d", l.input[:l.pos], l.input[l.pos:], l.pos)
	}
}

type item interface {
	result() bool
	add(item item)
}

type boolItem struct {
	item bool
}

func (b *boolItem) String() string {
	return fmt.Sprintf("boolItem{item=%v}", b.item)
}

func (b *boolItem) add(item item) {
	b.item = item.result()
}

func (b *boolItem) result() bool {
	return b.item
}

func (b *boolItem) getLast() item {
	return b
}

type notBoolItem struct {
	item item
}

func (b *notBoolItem) String() string {
	return fmt.Sprintf("notBoolItem{item=%v}", b.item)
}

func (b *notBoolItem) add(item item) {
	b.item = item
}

func (b *notBoolItem) result() bool {
	return !b.item.result()
}

type orGroupBoolItem struct {
	items []item
}

func (b *orGroupBoolItem) String() string {
	return fmt.Sprintf("orGroupBoolItem{items=%v}", b.items)
}

func (b *orGroupBoolItem) add(item item) {
	b.items = append(b.items, item)
}

func (b *orGroupBoolItem) result() bool {
	res := b.items[0].result()
	if res {
		return res
	}
	for i := 1; i < len(b.items); i++ {
		res = res || b.items[i].result()
		if res {
			return res
		}
	}
	return res
}

type andGroupBoolItem struct {
	items []item
}

func (b *andGroupBoolItem) String() string {
	return fmt.Sprintf("andGroupBoolItem{items=%v}", b.items)
}

func (b *andGroupBoolItem) add(item item) {
	b.items = append(b.items, item)
}

func (b andGroupBoolItem) result() bool {
	res := b.items[0].result()
	if !res {
		return res
	}
	for i := 1; i < len(b.items); i++ {
		res = res && b.items[i].result()
		if !res {
			return res
		}
	}
	return res
}

func parseBoolExpr(expression string) bool {
	lexer := newLexer(expression)
	for state := outerLex(lexer); state != nil; {
		state = state(lexer)
	}
	item := lexer.items.pop()
	return item.result()
}
