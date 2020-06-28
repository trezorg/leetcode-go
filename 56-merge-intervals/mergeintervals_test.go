package mergeintervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals1(t *testing.T) {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	output := [][]int{{1, 6}, {8, 10}, {15, 18}}
	assert.Equal(t, output, merge(input))
}

func TestMergeIntervals2(t *testing.T) {
	input := [][]int{{1, 3}, {8, 10}, {15, 18}}
	output := [][]int{{1, 3}, {8, 10}, {15, 18}}
	assert.Equal(t, output, merge(input))
}

func TestMergeIntervals3(t *testing.T) {
	input := [][]int{{1, 3}, {1, 3}, {1, 5}}
	output := [][]int{{1, 5}}
	assert.Equal(t, output, merge(input))
}

func TestMergeIntervals4(t *testing.T) {
	input := [][]int{{1, 4}, {4, 5}}
	output := [][]int{{1, 5}}
	assert.Equal(t, output, merge(input))
}

func TestMergeIntervals6(t *testing.T) {
	input := [][]int{}
	output := [][]int{}
	assert.Equal(t, output, merge(input))
}
