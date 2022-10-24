package water

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWater(t *testing.T) {

	data := []struct {
		data   []int
		result int
	}{
		// {[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		// {[]int{1, 1}, 1},
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%v", v.data), func(t *testing.T) {
			assert.Equal(t, v.result, maxArea(v.data))
		})
	}

}
