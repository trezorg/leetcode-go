package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var data = []struct {
	nums   []int
	target int
	result []int
}{
	{nums: []int{5, 7, 7, 8, 8, 10}, target: 8, result: []int{3, 4}},
	{nums: []int{5, 7, 7, 8, 8, 10}, target: 0, result: []int{-1, -1}},
	{nums: []int{}, target: 0, result: []int{-1, -1}},
	{nums: []int{5, 7, 7, 8, 8, 10}, target: 7, result: []int{1, 2}},
	{nums: []int{5, 7, 7, 8, 8, 10}, target: 10, result: []int{5, 5}},
	{nums: []int{5, 7, 7, 8, 8, 10, 10, 10}, target: 10, result: []int{5, 7}},
	{nums: []int{8, 8, 8, 8, 8, 8}, target: 8, result: []int{0, 5}},
}

func TestSearchRange(t *testing.T) {
	for _, v := range data {
		t.Run(fmt.Sprintf("%v-%d", v.nums, v.target), func(t *testing.T) {
			require.Equal(t, v.result, searchRange(v.nums, v.target))
		})
	}
}
func TestSearchRange2(t *testing.T) {
	for _, v := range data {
		t.Run(fmt.Sprintf("%v-%d", v.nums, v.target), func(t *testing.T) {
			require.Equal(t, v.result, searchRange2(v.nums, v.target))
		})
	}

}
