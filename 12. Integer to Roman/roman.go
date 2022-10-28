package roman

import (
	"bytes"
)

var digits []string = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var values []int = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	var res bytes.Buffer
	for i := 0; i < len(values); i++ {
		v := num / values[i]
		num = num % values[i]

		if v > 0 {
			for j := 0; j < v; j++ {
				res.WriteString(digits[i])
			}
		}

	}
	return res.String()
}
