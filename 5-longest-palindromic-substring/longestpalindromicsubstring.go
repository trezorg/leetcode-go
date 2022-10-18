package longestpalindromicsubstring

func longestSubPalindrome(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return i + 1, j
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	start, end := 0, 0

	middle := len(s) / 2

	for i, j := middle, middle+1; i >= 0 || j < len(s); i, j = i-1, j+1 {
		if (i+1)*2 < end-start && (len(s)-j+1)*2 < end-start {
			return s[start:end]
		}
		if i >= 0 {
			a, b := longestSubPalindrome(s, i, i)
			if b-a > end-start {
				start, end = a, b
			}
			a, b = longestSubPalindrome(s, i, i+1)
			if b-a > end-start {
				start, end = a, b
			}
		}
		if j < len(s) {
			a, b := longestSubPalindrome(s, j, j)
			if b-a > end-start {
				start, end = a, b
			}
			a, b = longestSubPalindrome(s, j, j+1)
			if b-a > end-start {
				start, end = a, b
			}
		}
	}

	return s[start:end]
}
