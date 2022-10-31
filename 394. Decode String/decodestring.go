package decodestring

import (
	"strings"
	"unicode"
	"strconv"
)

var (
	startGroup = '['
	endGroup = ']'
)

func decodeString(s string) string {
	var out strings.Builder
	idx := 0
	for idx < len(s) {
		bt := rune(s[idx])
		if !unicode.IsDigit(bt) && bt != startGroup {
			out.WriteRune(bt)
			idx++
		} else {
			var unwinded string
			unwinded, idx = unwindString(s, idx)
			out.WriteString(unwinded)
		}
	}
	return out.String()
}

func unwindString(s string, idx int) (string, int) {
	var number strings.Builder
	var str strings.Builder
	n := 1
	bt := rune(s[idx])
	for unicode.IsDigit(bt) {
		number.WriteRune(bt)
		idx++
		bt = rune(s[idx])
	}
	nStr := number.String()
	if len(nStr) > 0 {
		n, _ = strconv.Atoi(nStr)
	}
	for bt == startGroup {
		idx++
		bt = rune(s[idx])
	}
	for bt != ']' {
		if unicode.IsDigit(bt) {
			var unwinded string 
			unwinded, idx = unwindString(s, idx)
			str.WriteString(unwinded) 
		} else {
			str.WriteRune(bt)
			idx++
		}
		bt = rune(s[idx])
	}
	if bt == endGroup {
		repeat := str.String()
		for i := 0; i < n - 1; i++ {
			str.WriteString(repeat)
		} 
		idx++
	}
	return str.String(), idx
}
