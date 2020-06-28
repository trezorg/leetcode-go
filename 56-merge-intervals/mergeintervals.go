package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	j := 0
	for i := 1; i < len(intervals); i++ {
		if overlap(intervals[j], intervals[i]) {
			intervals[j][1] = max(intervals[j][1], intervals[i][1])
		} else {
			j++
			intervals[j] = intervals[i]
		}
	}
	return intervals[:min(j+1, len(intervals))]
}

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

func overlap(interval1 []int, interval2 []int) bool {
	return interval1[1] >= interval2[0] && interval1[0] <= interval2[1]
}
