package lispexpression

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	letOp         = "let"
	addOp         = "add"
	multOp        = "mult"
	digits        = "0123456789"
	signs         = "-+"
	alphabet      = "abcdefghijklmnopqrstuvwxyz"
	alphaDigits   = digits + alphabet
	digitsAndSign = digits + signs
	spaceSymbol   = rune(' ')
	groupStart    = rune('(')
	groupEnd      = rune(')')
	eof           = rune(0)
)

type state struct {
	values map[string]item
	prev   *state
}

func (s *state) find(name string) item {
	for s != nil {
		item, ok := s.values[name]
		if ok {
			return item
		}
		s = s.prev
	}
	return nil
}

func (s *state) addToState(name string, value item) *state {
	s.values[name] = value
	return s
}

func (s *state) getNewState() *state {
	state := &state{values: map[string]item{}, prev: s}
	return state
}

type lexer struct {
	input string
	start int
	pos   int
	width int
}

func newLexer(input string) *lexer {
	return &lexer{
		input: input,
	}
}

func newState(prev *state) *state {
	return &state{
		values: map[string]item{},
		prev:   prev,
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
	return fmt.Sprintf("%s - %s. Position: %d. Symbol :%c", l.input[:l.pos], l.input[l.pos:], l.pos, l.peek())
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

func (l *lexer) get() string {
	res := l.input[l.start:l.pos]
	l.forward()
	return res
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

type item interface {
	evaluate() int
	add(item item)
	getItems() []item
}

type valueItem struct {
	item int
}

func (v *valueItem) String() string {
	return fmt.Sprintf("valueItem{item=%d}", v.item)
}

func (v *valueItem) add(item item) {
	v.item = item.evaluate()
}

func (v *valueItem) evaluate() int {
	return v.item
}

func (v *valueItem) getItems() []item {
	return []item{v}
}

type variableItem struct {
	state *state
	name  string
}

func (v *variableItem) String() string {
	return fmt.Sprintf("variableItem{name=%s, value: %d}", v.name, v.evaluate())
}

func (v *variableItem) add(item item) {
	items := item.getItems()
	for i := 0; i < len(items); i++ {
		val, ok := items[i].(*variableItem)
		if !ok {
			continue
		}
		if val.name == v.name {
			val.state = val.state.getNewState().addToState(v.name, &valueItem{item: val.evaluate()})
		}
	}
	v.state.addToState(v.name, item)
}

func (v *variableItem) evaluate() int {
	item := v.state.find(v.name)
	if item == nil {
		fmt.Printf("Cannot find state for valiable: %s, %v\n", v.name, v.state)
	}
	return item.evaluate()
}

func (v *variableItem) getItems() []item {
	return []item{v}
}

type addItem struct {
	items []item
}

func (a *addItem) String() string {
	return fmt.Sprintf("addItem{items=%s}", a.items)
}

func (a *addItem) add(item item) {
	a.items = append(a.items, item)
}

func (a *addItem) evaluate() int {
	if len(a.items) == 0 {
		return 0
	}
	res := a.items[0].evaluate()
	for i := 1; i < len(a.items); i++ {
		res = res + a.items[i].evaluate()
	}
	return res
}

func (a *addItem) getItems() []item {
	return a.items
}

type multItem struct {
	items []item
}

func (m *multItem) String() string {
	return fmt.Sprintf("multItem{items=%s}", m.items)
}

func (m *multItem) add(item item) {
	m.items = append(m.items, item)
}

func (m *multItem) evaluate() int {
	if len(m.items) == 0 {
		return 0
	}
	res := m.items[0].evaluate()
	for i := 1; i < len(m.items); i++ {
		res = res * m.items[i].evaluate()
	}
	return res
}

func (m *multItem) getItems() []item {
	return m.items
}

func groupLexer(l *lexer, state *state) (item, error) {
	l.swallow(spaceSymbol)
	switch l.next() {
	case groupStart:
		l.forward()
		l.swallow(spaceSymbol)
		l.acceptTill(spaceSymbol)
		switch l.get() {
		case multOp:
			l.forward()
			return groupItemLexer(l, state, &multItem{})
		case addOp:
			l.forward()
			return groupItemLexer(l, state, &addItem{})
		case letOp:
			l.forward()
			return letItemLexer(l, state.getNewState())
		default:
			return nil, fmt.Errorf("group lexer. bad syntax: %s - %s. Position: %d. Symbol: %c", l.input[:l.pos], l.input[l.pos:], l.pos, l.peek())
		}
	default:
		return nil, fmt.Errorf("group lexer. bad syntax: %s - %s. Position: %d. Symbol: %c", l.input[:l.pos], l.input[l.pos:], l.pos, l.peek())
	}
}

func groupItemLexer(l *lexer, state *state, item item) (item, error) {
	l.swallow(spaceSymbol)
	for s := l.peek(); s != groupEnd; {
		if l.accept(alphabet) {
			// variable
			l.acceptRun(alphaDigits)
			item.add(&variableItem{name: l.get(), state: state})
		} else if l.accept(digitsAndSign) {
			l.acceptRun(digits)
			newItemValue, _ := strconv.Atoi(l.get())
			item.add(&valueItem{item: newItemValue})
		} else if l.acceptRune(groupStart) {
			l.backup()
			newItem, err := groupLexer(l, state)
			if err != nil {
				return nil, err
			}
			item.add(newItem)
		} else if s == eof {
			break
		} else {
			return nil, fmt.Errorf("group item lexer. bad syntax: %s - %s. Position: %d. Symbol: %c", l.input[:l.pos], l.input[l.pos:], l.pos, l.peek())
		}
		l.swallow(spaceSymbol)
		s = l.peek()
	}
	l.swallow(spaceSymbol)
	l.swallowOne(groupEnd)
	l.swallow(spaceSymbol)
	return item, nil
}

func letItemLexer(l *lexer, state *state) (item, error) {
	l.swallow(spaceSymbol)
	state = state.getNewState()
	var result item
	var err error
	elmCount := 0
	for s := l.peek(); s != groupEnd; {
		if l.acceptRune(groupStart) {
			l.backup()
			result, err = groupLexer(l, state)
			if err != nil {
				return nil, err
			}
			elmCount++
		} else if l.accept(digitsAndSign) {
			l.acceptRun(digits)
			itemValue, _ := strconv.Atoi(l.get())
			result = &valueItem{item: itemValue}
			elmCount++
		} else if l.accept(alphabet) {
			// variable
			l.acceptRun(alphaDigits)
			result = &variableItem{name: l.get(), state: state}
			l.swallow(spaceSymbol)
			elmCount++
			if l.accept(digitsAndSign) {
				l.acceptRun(digits)
				itemValue, _ := strconv.Atoi(l.get())
				result.add(&valueItem{item: itemValue})
				elmCount++
			} else if l.accept(alphabet) {
				l.acceptRun(alphaDigits)
				variableResut := &variableItem{name: l.get(), state: state}
				result.add(variableResut)
				l.swallow(spaceSymbol)
				elmCount++
			} else if l.acceptRune(groupStart) {
				l.backup()
				groupResult, err := groupLexer(l, state)
				if err != nil {
					return nil, err
				}
				result.add(groupResult)
				elmCount++
			}
		} else if s == eof {
			break
		} else {
			return nil, fmt.Errorf("let item lexer. bad syntax: %s - %s. Position: %d", l.input[:l.pos], l.input[l.pos:], l.pos)
		}
		l.swallow(spaceSymbol)
		s = l.peek()
	}
	l.swallow(spaceSymbol)
	l.swallowOne(groupEnd)
	l.swallow(spaceSymbol)
	if elmCount%2 == 0 {
		return nil, fmt.Errorf("let item lexer. bad syntax: %s - %s. Position: %d. No output variable", l.input[:l.pos], l.input[l.pos:], l.pos)
	}
	return result, nil
}

func evaluate(expression string) int {
	lexer := newLexer(expression)
	state := newState(nil)
	item, err := groupLexer(lexer, state)
	if err != nil {
		panic(err)
	}
	return item.evaluate()
}
