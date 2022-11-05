package main

const size = 'z' - 'a' + 1

func checkInclusion(s1 string, s2 string) bool {
	n, sn := len(s1), len(s2)
	if sn < n {
		return false
	}
	subStringCounts := [size]int{}
	stringCounts := [size]int{}

	for i := 0; i < n; i++ {
		subStringCounts[s1[i]-'a']++
		stringCounts[s2[i]-'a']++
	}

	if subStringCounts == stringCounts {
		return true
	}

	for i := 1; i < sn-n+1; i++ {
		stringCounts[s2[i-1]-'a']--
		stringCounts[s2[i+n-1]-'a']++
		if subStringCounts == stringCounts {
			return true
		}
	}

	return false
}
