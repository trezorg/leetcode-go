package foursum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {

	data := []struct {
		data   []int
		target int
		result [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{[]int{2, 2, 2, 2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{[]int{-1, 0, 1, 2, -1, -4}, -1, [][]int{{-4, 0, 1, 2}, {-1, -1, 0, 1}}},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%v", v.data), func(t *testing.T) {
			assert.Equal(t, v.result, fourSum(v.data, v.target))
		})
	}

}
