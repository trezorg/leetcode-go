package reverse

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {

	data := []struct {
		data   int
		result int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{math.MaxInt32, 0},
		{-math.MaxInt32, 0},
		{1463847412, 2147483641},
		{-1463847412, -2147483641},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%v", v.data), func(t *testing.T) {
			assert.Equal(t, v.result, reverse(v.data))
		})
	}

}
