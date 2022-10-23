package threesum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {

	data := []struct {
		data   []int
		result [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 0}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{0, 0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{[]int{1, 1, -2}, [][]int{{-2, 1, 1}}},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%v", v.data), func(t *testing.T) {
			assert.Equal(t, v.result, threeSum(v.data))
		})
	}

}
