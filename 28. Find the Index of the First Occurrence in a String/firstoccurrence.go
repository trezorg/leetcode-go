package main

func comp(a, b []byte) int {
	var j int
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			// if more then one first symbols from right are equal, shift 1
			// here should be optimization for Boyer–Moore algorithm
			if j > 0 {
				return 1
			}
			// if first is not equal
			break
		}
		j++
	}
	if j == len(a) {
		return -1
	}
	// find first symbol from right and calculate shift
	f := a[len(a)-1]
	for i, j := len(b)-1, 0; i >= 0; i, j = i-1, j+1 {
		if f == b[i] {
			return j
		}
	}
	// shift all string
	return len(a)
}

func strStr(haystack string, needle string) int {
	sn, n := len(haystack), len(needle)
	if sn < n {
		return -1
	}
	a, b := []byte(haystack), []byte(needle)
	for i := 0; i < sn-n+1; {
		var idx int
		idx = comp(a[i:i+n], b)
		if idx == -1 {
			return i
		}
		i += idx
	}

	return -1

}
