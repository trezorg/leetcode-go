package water

func maxArea(height []int) int {
	area := 0

	i, j := 0, len(height)-1
	for i < j {
		a, b := height[i], height[j]
		var mul int
		if a < b {
			mul = a
		} else {
			mul = b
		}

		temp := (j - i) * mul
		if area < temp {
			area = temp
		}
		if a < b {
			i++
		} else {
			j--
		}
	}
	return area
}
