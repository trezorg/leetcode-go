package main

func checkInclusion(s1 string, s2 string) bool {
	n, sn := len(s1), len(s2)
	if sn < n {
		return false
	}
	subStringCounts := [127]int{}
	stringCounts := [127]int{}

	for i := 0; i < n; i++ {
		subStringCounts[s1[i]-'0']++
		stringCounts[s2[i]-'0']++
	}

	if subStringCounts == stringCounts {
		return true
	}

	for i := 1; i < sn-n+1; i++ {
		stringCounts[s2[i-1]-'0']--
		stringCounts[s2[i+n-1]-'0']++
		if subStringCounts == stringCounts {
			return true
		}
	}

	return false
}
