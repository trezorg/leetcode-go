package trapwater

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// trap. Calculate area of water
// find left and right max for each index
// calculate min(leftMax, rightMax) * height
func trap(height []int) int {

	type leftRight [2]int

	maxLeftRight := make([]leftRight, len(height))

	for i, j := 1, len(height)-2; i < len(height) && j >= 0; i, j = i+1, j-1 {
		maxLeftRight[i][0] = max(height[i-1], maxLeftRight[i-1][0])
		maxLeftRight[j][1] = max(height[j+1], maxLeftRight[j+1][1])
	}

	area := 0

	for i := 0; i < len(height)-1; i++ {
		area += max(0, min(maxLeftRight[i][0], maxLeftRight[i][1])-height[i])
	}

	return area

}

func trap2(height []int) int {

	mx, mi, area := 0, 0, 0

	for i, j := 0, len(height)-1; i <= j; {
		mi = min(height[i], height[j])
		mx = max(mi, mx)
		area += mx - mi
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return area

}
