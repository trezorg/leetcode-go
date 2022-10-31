package booleanexpression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolExpression1(t *testing.T) {
	s := "!(f)"
	assert.True(t, parseBoolExpr(s))
}

func TestBoolExpression2(t *testing.T) {
	s := "!(t)"
	assert.False(t, parseBoolExpr(s))
}

func TestBoolExpression3(t *testing.T) {
	s := "|(f,t)"
	assert.True(t, parseBoolExpr(s))
}

func TestBoolExpression4(t *testing.T) {
	s := "|(f,f)"
	assert.False(t, parseBoolExpr(s))
}

func TestBoolExpression5(t *testing.T) {
	s := "&(t,f)"
	assert.False(t, parseBoolExpr(s))
}

func TestBoolExpression6(t *testing.T) {
	s := "&(f,f)"
	assert.False(t, parseBoolExpr(s))
}

func TestBoolExpression7(t *testing.T) {
	s := "&(t,t)"
	assert.True(t, parseBoolExpr(s))
}

func TestBoolExpression8(t *testing.T) {
	s := "|(&(t,f,t),!(t))"
	assert.False(t, parseBoolExpr(s))
}

func TestBoolExpression9(t *testing.T) {
	s := "|(&(t,f,t),!(f))"
	assert.True(t, parseBoolExpr(s))
}

func TestBoolExpression10(t *testing.T) {
	s := "!(f)"
	assert.True(t, parseBoolExpr(s))
}

func TestBoolExpression11(t *testing.T) {
	s := "|(f,!(f))"
	assert.True(t, parseBoolExpr(s))
}

func TestBoolExpression12(t *testing.T) {
	s := "|(f,!(t))"
	assert.False(t, parseBoolExpr(s))
}
