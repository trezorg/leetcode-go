package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var data = []struct {
	dividend int
	divisor  int
	result   int
}{
	{dividend: 10, divisor: 3, result: 3},
	{dividend: 7, divisor: -3, result: -2},
	{dividend: 0, divisor: -3, result: 0},
	{dividend: 120, divisor: 4, result: 30},
	{dividend: -2147483648, divisor: -1, result: 2147483647},
	{dividend: 2147483648, divisor: 1, result: 2147483647},
	{dividend: -1, divisor: 1, result: -1},
	{dividend: 1, divisor: 1, result: 1},
	{dividend: 15, divisor: 5, result: 3},
	{dividend: -2147483648, divisor: -2147483648, result: 1},
}

func TestStrStr(t *testing.T) {
	for _, v := range data {
		t.Run(fmt.Sprintf("%d-%d", v.dividend, v.divisor), func(t *testing.T) {
			require.Equal(t, v.result, divide(v.dividend, v.divisor))
		})
	}

}
