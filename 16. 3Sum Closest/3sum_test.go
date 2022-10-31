package threesum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {

	data := []struct {
		data   []int
		target int
		result int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 0, 0}, 1, 0},
		{[]int{1, 1, 1}, 3, 3},
		{[]int{1, 2, 3}, 3, 6},
		{[]int{1, 2, 3, 4}, 3, 6},
		{[]int{-1, -2, 30, 4}, 1, 1},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%v", v.data), func(t *testing.T) {
			assert.Equal(t, v.result, threeSumClosest(v.data, v.target))
		})
	}

}
