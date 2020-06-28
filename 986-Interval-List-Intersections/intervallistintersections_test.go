package intervallistintersections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals1(t *testing.T) {
	input1 := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	input2 := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	output := [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}}
	assert.Equal(t, output, intervalIntersection(input1, input2))
}

func TestMergeIntervals2(t *testing.T) {
	input1 := [][]int{{0, 2}, {5, 10}, {13, 23}}
	input2 := [][]int{{25, 26}}
	output := [][]int{}
	assert.Equal(t, output, intervalIntersection(input1, input2))
}

func TestMergeIntervals3(t *testing.T) {
	input1 := [][]int{{0, 2}, {5, 10}, {13, 23}}
	input2 := [][]int{{6, 7}}
	output := [][]int{{6, 7}}
	assert.Equal(t, output, intervalIntersection(input1, input2))
}

func TestMergeIntervals4(t *testing.T) {
	input1 := [][]int{}
	input2 := [][]int{{6, 7}}
	output := [][]int{}
	assert.Equal(t, output, intervalIntersection(input1, input2))
}

func TestMergeIntervals5(t *testing.T) {
	input1 := [][]int{{0, 2}, {5, 10}}
	input2 := [][]int{{1, 6}}
	output := [][]int{{1, 2}, {5, 6}}
	assert.Equal(t, output, intervalIntersection(input1, input2))
}
