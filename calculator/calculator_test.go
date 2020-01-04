package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse1(t *testing.T) {
	s := "1 + 2 + 3"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 6, node.value())
}

func TestParse2(t *testing.T) {
	s := "1 + 2 - 3"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 0, node.value())
}

func TestParse3(t *testing.T) {
	s := "(1 + 2) + 3"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 6, node.value())
}

func TestParse4(t *testing.T) {
	s := "(1 + 2) + (3"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 6, node.value())
}

func TestParse5(t *testing.T) {
	s := "(1 + 2)   +      (3 + 2)"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 8, node.value())
}

func TestParse6(t *testing.T) {
	s := "(1 + 2)   *      (3 + 2)"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 15, node.value())
}

func TestParse7(t *testing.T) {
	s := "1 + 2   *      (3 + 2 * (11 + 7))"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 79, node.value())
}

func TestParse8(t *testing.T) {
	s := "(1 + 2)   *      (3 + 2 * (11 + 7))"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 117, node.value())
}

func TestParse9(t *testing.T) {
	s := "(1 + 9)   % (1 + 2 + 3 - 2 + 4 * 1 - 2 - 1)"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 0, node.value())
}

func TestParse10(t *testing.T) {
	s := "(1 + 9)  / (1 + 2 + 3 - 2 + 4 * 1 - 2 - 1)"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 2, node.value())
}

func TestParse11(t *testing.T) {
	s := "(1 1 1)"
	_, _, err := parse(s)
	require.Error(t, err)
}

func TestParse12(t *testing.T) {
	s := "(1 + 1  / 1 + 1)"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 3, node.value())
}

func TestParse13(t *testing.T) {
	s := "(1 + 1  / 2 - 1)"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 0, node.value())
}

func TestParse14(t *testing.T) {
	s := "1024 / 2 / 2 / 2 / 2"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 64, node.value())
}

func TestParse15(t *testing.T) {
	s := "1024 / 2 / (2 * 2 * 2)"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 64, node.value())
}

func TestParse16(t *testing.T) {
	s := "1024 / 2 * (2 - 2 + 2)"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 1024, node.value())
}

func TestParse17(t *testing.T) {
	s := "1024 / 2 * ((2 - 1) / 2)"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 0, node.value())
}

func TestParse18(t *testing.T) {
	s := "1024 + 2 * ((((2 + 1) / 2 + 12)))))))))"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 1050, node.value())
}

func TestParse19(t *testing.T) {
	s := "(1+(4+5+2)-3)+(6+8)"
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 23, node.value())
}

func TestParse20(t *testing.T) {
	s := " 2-1 + 2 "
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 3, node.value())
}

func TestParse21(t *testing.T) {
	s := " 2-1 + 2 * 100 "
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 201, node.value())
}

func TestParse22(t *testing.T) {
	s := " 3+5 / 2 "
	node, _, err := parse(s)
	require.NoError(t, err)
	assert.Equal(t, 5, node.value())
}


func BenchmarkTestParse(b *testing.B) {
	s := "(1+(4+5+2)-3)+(6+8)"
	for i := 0; i <= b.N; i++ {
		node, _,  _ := parse(s)
		assert.Equal(b, 23, node.value())
	}
}
