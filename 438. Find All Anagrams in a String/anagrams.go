package main

const size = 'z' - 'a' + 1

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	n, sn := len(p), len(s)
	if sn < n {
		return res
	}
	subStringCounts := [size]int{}
	stringCounts := [size]int{}

	for i := 0; i < n; i++ {
		subStringCounts[p[i]-'a']++
		stringCounts[s[i]-'a']++
	}

	if subStringCounts == stringCounts {
		res = append(res, 0)
	}

	for i := 1; i < sn-n+1; i++ {
		stringCounts[s[i-1]-'a']--
		stringCounts[s[i+n-1]-'a']++
		if subStringCounts == stringCounts {
			res = append(res, i)
		}
	}

	return res
}
