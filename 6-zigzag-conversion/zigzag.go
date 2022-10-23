package zigzag

import "unsafe"

func convert(s string, numRows int) string {
	if len(s) < numRows || numRows <= 1 {
		return s
	}
	sl := make([][]byte, 0, numRows)
	sl = append(sl, make([]byte, 0, len(s)))
	for i := 1; i < numRows; i++ {
		sl = append(sl, make([]byte, 0, len(s)/numRows))
	}

	step, n := 1, 0
	for i := 0; i < len(s); i++ {
		if n == 0 {
			step = 1
		}
		if n == numRows-1 {
			step = -1
		}
		sl[n] = append(sl[n], s[i])
		n += step
	}

	for i := 1; i < numRows; i++ {
		sl[0] = append(sl[0], sl[i]...)
	}

	return *(*string)(unsafe.Pointer(&sl[0]))

}
