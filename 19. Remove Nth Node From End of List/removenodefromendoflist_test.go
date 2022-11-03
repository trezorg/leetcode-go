package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromList(t *testing.T) {

	data := []struct {
		data   []int
		n      int
		result []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{2}},
		{[]int{}, 2, []int{}},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%v", v.data), func(t *testing.T) {
			head := newList(v.data)
			assert.Equal(t, v.result, removeNthFromEnd(head, v.n).Slice())
		})
	}

}
