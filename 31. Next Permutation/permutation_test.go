package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {

	data := []struct {
		data   []int
		result []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%v", v.data), func(t *testing.T) {
			nextPermutation(v.data)
			assert.Equal(t, v.result, v.data)
		})
	}

}
