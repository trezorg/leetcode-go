package lispexpression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpression1(t *testing.T) {
	s := "(let x 2 (mult x (let x 3 y 4 (add x y))))"
	assert.Equal(t, 14, evaluate(s))
}

func TestExpression2(t *testing.T) {
	s := "(add 1 2)"
	assert.Equal(t, 3, evaluate(s))
}

func TestExpression3(t *testing.T) {
	s := "(let x 12 (add x 1))"
	assert.Equal(t, 13, evaluate(s))
}

func TestExpression4(t *testing.T) {
	s := "(let x 12 (mult x 2))"
	assert.Equal(t, 24, evaluate(s))
}

func TestExpression5(t *testing.T) {
	s := "(add (let x 12 (add x 1)) 1 2 3)"
	assert.Equal(t, 19, evaluate(s))
}

func TestExpression6(t *testing.T) {
	s := "(add (let x 12 (add x 1)) (let x 13 (add x 1)))"
	assert.Equal(t, 27, evaluate(s))
}

func TestExpression7(t *testing.T) {
	s := "(add 1 2)"
	assert.Equal(t, 3, evaluate(s))
}

func TestExpression8(t *testing.T) {
	s := "(mult 3 (add 2 3))"
	assert.Equal(t, 15, evaluate(s))
}
func TestExpression9(t *testing.T) {
	s := "(let x 2 (mult x 5))"
	assert.Equal(t, 10, evaluate(s))
}
func TestExpression10(t *testing.T) {
	s := "(let x 2 (mult x (let x 3 y 4 (add x y))))"
	assert.Equal(t, 14, evaluate(s))
}
func TestExpression11(t *testing.T) {
	s := "(let x 3 x 2 x)"
	assert.Equal(t, 2, evaluate(s))
}
func TestExpression12(t *testing.T) {
	s := "(let x 1 y 2 x (add x y) (add x y))"
	assert.Equal(t, 5, evaluate(s))
}

func TestExpression13(t *testing.T) {
	s := "(let x 2 (add (let x 3 (let x 4 x)) x))"
	assert.Equal(t, 6, evaluate(s))
}

func TestExpression14(t *testing.T) {
	s := "(let a1 3 b2 (add a1 1) b2)"
	assert.Equal(t, 4, evaluate(s))
}

func TestExpression15(t *testing.T) {
	s := "(let x 7 -12)"
	assert.Equal(t, -12, evaluate(s))
}

func TestExpression16(t *testing.T) {
	s := "(let x 7 y (add -12 x) (add x y))"
	assert.Equal(t, 2, evaluate(s))
}

func TestExpression17(t *testing.T) {
	s := "(let x 7 y (add -12 -27) (add x y))"
	assert.Equal(t, -32, evaluate(s))
}

func TestExpression18(t *testing.T) {
	s := "(let x -2 y x y)"
	assert.Equal(t, -2, evaluate(s))
}

func TestExpression19(t *testing.T) {
	s := "(let x -2 x x x)"
	assert.Equal(t, -2, evaluate(s))
}

func TestExpression20(t *testing.T) {
	s := "(let x -2 z x y z (add y z x))"
	assert.Equal(t, -6, evaluate(s))
}

func TestExpression21(t *testing.T) {
	s := "(let x -2 z x a z t (add y z x) y 2 t)"
	assert.Equal(t, -2, evaluate(s))
}

func TestExpression22(t *testing.T) {
	s := "(let x -2 z (add x 1) y 4 (add y z))"
	assert.Equal(t, 3, evaluate(s))
}

func TestLexer1(t *testing.T) {
	s := "xxxxxyyyyyysssss"
	lexer := newLexer(s)
	lexer.acceptRun("x")
	assert.Equal(t, "xxxxx", lexer.get())
	lexer.acceptRun("y")
	assert.Equal(t, "yyyyyy", lexer.get())
}

func TestLexer2(t *testing.T) {
	s := "xxxxxyyyyyysssss"
	lexer := newLexer(s)
	lexer.acceptRun("xy")
	assert.Equal(t, "xxxxxyyyyyy", lexer.get())
}

func TestLexer3(t *testing.T) {
	s := "xxxxxyyyyyysssss"
	lexer := newLexer(s)
	lexer.acceptRun("xys")
	assert.Equal(t, "xxxxxyyyyyysssss", lexer.get())
}

func BenchmarkExpression(b *testing.B) {
	s := "(let x 2 (mult x (let x 3 y 4 (add x y))))"
	for i := 0; i < b.N; i++ {
		assert.Equal(b, 14, evaluate(s))
	}
}
