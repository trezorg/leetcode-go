package intervallistintersections

func intervalIntersection(A [][]int, B [][]int) [][]int {
	result := make([][]int, 0)
	if len(A) == 0 || len(B) == 0 {
		return result
	}
	startAIdx, startBIdx, endAIdx, endBIdx := 0, 0, 0, 0
	counter := 0
	var prevValue int

	nextStartValue := func() int {
		startValue := -1
		if startAIdx < len(A) && startBIdx < len(B) {
			if A[startAIdx][0] < B[startBIdx][0] {
				startValue = A[startAIdx][0]
				startAIdx++
			} else {
				startValue = B[startBIdx][0]
				startBIdx++
			}
		} else if startAIdx < len(A) {
			startValue = A[startAIdx][0]
			startAIdx++
		} else if startBIdx < len(B) {
			startValue = B[startBIdx][0]
			startBIdx++
		}
		return startValue
	}

	nextEndValue := func() int {
		endValue := -1
		if endAIdx < len(A) && endBIdx < len(B) {
			if A[endAIdx][1] < B[endBIdx][1] {
				endValue = A[endAIdx][1]
				endAIdx++
			} else {
				endValue = B[endBIdx][1]
				endBIdx++
			}
		} else if endAIdx < len(A) {
			endValue = A[endAIdx][1]
			endAIdx++
		} else if endBIdx < len(B) {
			endValue = B[endBIdx][1]
			endBIdx++
		}
		return endValue
	}

	endValue := nextEndValue()
	startValue := nextStartValue()

	for endValue != -1 && startValue != -1 {
		if startValue <= endValue {
			counter++
			if counter == 2 {
				prevValue = startValue
			}
			startValue = nextStartValue()
		} else {
			if counter == 2 {
				result = append(result, []int{prevValue, endValue})
			}
			counter--
			endValue = nextEndValue()
		}
	}
	for endValue != -1 {
		if counter == 2 {
			result = append(result, []int{prevValue, endValue})
		}
		counter--
		endValue = nextEndValue()
	}

	return result
}
