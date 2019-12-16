package regexp

import (
	"fmt"
	"math"
)

const (
	dot            = byte('.')
	star           = byte('*')
	infiniteRepeat = math.MaxInt32
	anySymbol      = byte(0)
	blankSymbol    = byte(0)
)

var (
	alphabet    = "abcdefghijklmnopqrstuvwxyz"
	alphabetMap = func() map[byte]bool {
		res := map[byte]bool{}
		for _, s := range alphabet {
			res[byte(s)] = true
		}
		return res

	}()
	dotToken             = SingleSymbolToken{symbol: dot}
	starToken            = SingleSymbolToken{symbol: star}
	alphabetSymbolToken  = OneOfSymbolToken{symbols: alphabetMap}
	greedyAnySymbolToken = GroupSymbolToken{tokens: []Token{dotToken, starToken}}
	greedySymbolToken    = GroupSymbolToken{tokens: []Token{alphabetSymbolToken, starToken}}
	allowedTokens        = []Token{greedyAnySymbolToken, greedySymbolToken, dotToken, alphabetSymbolToken}
)

// Token interface
type Token interface {
	read(data []byte) int
}

// OneOfSymbolToken for one of possible tokens
type OneOfSymbolToken struct {
	symbols map[byte]bool
}

func (t OneOfSymbolToken) read(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	_, ok := t.symbols[data[0]]
	if !ok {
		return 0
	}
	return 1
}

// SingleSymbolToken one token
type SingleSymbolToken struct {
	symbol byte
}

func (t SingleSymbolToken) read(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	if t.symbol == data[0] {
		return 1
	}
	return 0
}

// GroupSymbolToken group of tokens
type GroupSymbolToken struct {
	tokens []Token
}

func (t GroupSymbolToken) read(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	read := 0
	for _, token := range t.tokens {
		n := token.read(data)
		if n == 0 {
			return read
		}
		read += n
		data = data[n:]
	}
	return read
}

// Pos struct
type Pos struct {
	data   int
	regexp int
}

// Node regex representation
type Node struct {
	symbol    byte
	repeatMin int
	repeatMax int
}

func (n Node) equal(b byte) bool {
	if n.repeatMin == 0 && b == blankSymbol {
		return true
	}
	if b == blankSymbol {
		return false
	}
	if n.symbol == anySymbol {
		return true
	}
	return n.symbol == b
}

// Regexp regex representation
type Regexp struct {
	nodes []Node
}

func (r Regexp) isEmpty() bool {
	return len(r.nodes) == 0
}

func (r Regexp) checkPos(data []byte, pos Pos) bool {

	node := r.nodes[pos.regexp]
	b := data[pos.data]

	if node.equal(b) && pos.regexp == len(r.nodes)-1 && (pos.data == len(data)-1 || pos.data == len(data)-2 && data[len(data)-1] == blankSymbol) {
		return true
	}

	return false
}

func (r Regexp) nextSteps(data []byte, pos Pos) []Pos {

	node := r.nodes[pos.regexp]
	b := data[pos.data]
	var steps []Pos

	if !node.equal(b) {
		if node.repeatMin == 0 {
			steps := append(steps, Pos{pos.data, pos.regexp + 1})
			return steps
		}
		if node.repeatMin > 0 {
			return steps
		}
	}
	if node.repeatMin == 0 {
		steps = append(steps, Pos{pos.data, pos.regexp + 1})
	}
	steps = append(steps, Pos{pos.data + 1, pos.regexp + 1})
	if node.repeatMax > 1 {
		steps = append(steps, Pos{pos.data + 1, pos.regexp})
	}
	return steps
}

func (r Regexp) len() int {
	return len(r.nodes)
}

func nodeFromBytes(data []byte) (*Node, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("cannot parse blank data")
	}
	if len(data) == 1 {
		if data[0] == dot {
			return &Node{symbol: anySymbol, repeatMin: 1, repeatMax: 1}, nil
		}
		return &Node{symbol: data[0], repeatMin: 1, repeatMax: 1}, nil
	}
	if data[0] == dot {
		return &Node{symbol: anySymbol, repeatMin: 0, repeatMax: infiniteRepeat}, nil
	}
	return &Node{symbol: data[0], repeatMin: 0, repeatMax: infiniteRepeat}, nil
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

func mergeNodes(nodes []Node) []Node {
	if len(nodes) < 2 {
		return nodes
	}
	i, j, end := 0, 1, len(nodes)-1
	for j < len(nodes) {
		fNode := nodes[i]
		sNode := nodes[j]
		if fNode.symbol == sNode.symbol && fNode.repeatMin == 0 && sNode.repeatMin == 0 {
			nodes[i].repeatMin = 0
			nodes[i].repeatMax = infiniteRepeat
			j++
		} else {
			if j-i > 1 {
				nodes[i+1] = sNode
				end = i + 1
			}
			i++
			j++
		}
	}
	return nodes[:end+1]

}

func parseRegexp(data string) (*Regexp, error) {
	parts, err := parseTokens(data)
	if err != nil {
		return nil, err
	}
	var nodes []Node
	for _, part := range parts {
		node, err := nodeFromBytes(part)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, *node)
	}

	return &Regexp{nodes: mergeNodes(nodes)}, nil
}

func isMatch(s string, p string) bool {

	regexp, err := parseRegexp(p)
	if err != nil {
		return false
	}
	if len(s) == 0 && regexp.isEmpty() {
		return true
	}
	if regexp.isEmpty() {
		return false
	}
	regexLen := regexp.len()
	data := []byte(s)
	data = append(data, blankSymbol)
	stack := []Pos{{0, 0}}

	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		pos := stack[lastIdx]
		stack = stack[:lastIdx]
		if regexp.checkPos(data, pos) {
			return true
		}
		positions := regexp.nextSteps(data, pos)
		for _, p := range positions {
			if p.regexp < regexLen && p.data < len(data) {
				stack = append(stack, p)
			}
		}
	}

	return false

}
