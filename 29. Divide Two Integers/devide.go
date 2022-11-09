package main

import "math"

func divide(dividend int, divisor int) int {

	d := 0
	temp := 0

	var multi = 1
	if dividend < 0 {
		dividend = -dividend
		multi *= -1
	}
	if divisor < 0 {
		divisor = -divisor
		multi *= -1
	}

	n := int(math.Log2(float64(dividend)) + 1)

	for i := n; i >= 0; i-- {
		shift := divisor << i
		// overflow
		if shift < 0 {
			continue
		}
		if temp+shift <= dividend {
			temp += shift
			d |= 1 << i
		}
		if multi < 0 && d > math.MaxInt32+1 {
			return math.MinInt32
		}
		if multi > 0 && d > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return d * multi
}
