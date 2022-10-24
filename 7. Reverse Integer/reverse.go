package reverse

import (
	"math"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if res > 0 && res*10+pop > math.MaxInt32 {
			return 0
		}
		if res < 0 && res*10+pop < math.MinInt32 {
			return 0
		}
		res = res*10 + pop
	}
	return res
}
