package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var pivotData = []struct {
	nums   []int
	result int
}{
	{nums: []int{4, 5, 6, 7, 0, 1, 2}, result: 4},
	{nums: []int{5, 6, 7, 0, 1, 2, 4}, result: 3},
	{nums: []int{6, 7, 0, 1, 2, 4, 5}, result: 2},
	{nums: []int{7, 0, 1, 2, 4, 5, 6}, result: 1},
	{nums: []int{1, 2, 4, 5, 6, 7, 0}, result: 6},
	{nums: []int{0, 0, 0, 0, 0}, result: 0},
	{nums: []int{1, 1, 1, 1, 1}, result: 0},
}

var searchData = []struct {
	nums   []int
	target int
	result int
}{
	{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 7, result: 3},
	{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 1, result: 5},
	{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 2, result: 6},
	{nums: []int{1, 3, 6, 7, 9, 11, 12}, target: 11, result: 5},
	{nums: []int{}, target: 11, result: -1},
	{nums: []int{6, 7, 0, 1, 2, 4, 5}, target: 7, result: 1},
	{nums: []int{12}, target: 11, result: -1},
	{nums: []int{11}, target: 11, result: 0},
}

func TestSearchPivot(t *testing.T) {
	for _, v := range pivotData {
		t.Run(fmt.Sprintf("%v-%d", v.nums, v.result), func(t *testing.T) {
			require.Equal(t, v.result, searchPivot(v.nums))
		})
	}
}

func TestBinarySearch(t *testing.T) {
	for _, v := range searchData {
		t.Run(fmt.Sprintf("%v-%d-%d", v.nums, v.target, v.result), func(t *testing.T) {
			require.Equal(t, v.result, search(v.nums, v.target))
		})
	}
}
