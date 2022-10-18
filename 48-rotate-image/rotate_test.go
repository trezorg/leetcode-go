package rotate

import (
	"testing"

	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {

	type value struct {
		data   [][]int
		result [][]int
	}

	data := []value{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, [][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}}},
		{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%v", v.data), func(t *testing.T) {
			rotate(v.data)
			assert.Equal(t, v.result, v.data)
		})
	}

}
