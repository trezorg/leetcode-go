package phoneletters

import (
	"unsafe"
)

func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	table := [][]byte{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}
	l := 1

	for i := range digits {
		n := digits[i] - '0'
		l *= len(table[n])
	}

	results := make([][]byte, l)

	for i := 0; i < l; i++ {
		results[i] = make([]byte, len(digits))
	}

	period := l
	/*  Cartesian product
	Example abc, def
	a a a b b b c c c
	ad ae af bd be bf cd ce cf
	*/
	for i := 0; i < len(digits); i++ {
		n := digits[i] - '0'
		t := table[n]
		period /= len(t)
		for j := 0; j < l; j++ {
			results[j][i] = t[(j/period)%len(t)]
		}
	}

	out := make([]string, l)
	for i := range out {
		out[i] = *(*string)(unsafe.Pointer(&results[i]))
	}
	return out

}
