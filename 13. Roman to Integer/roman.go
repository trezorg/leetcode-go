package roman

var digits []string = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var values []int = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func romanToInt(s string) int {
	var res int
	for i := 0; i < len(digits); i++ {
		if len(s) == 0 {
			break
		}
		l := len(digits[i])
		for len(s) > 0 && len(s) >= l && s[:l] == digits[i] {
			res += values[i]
			s = s[l:]
		}
	}
	return res
}
