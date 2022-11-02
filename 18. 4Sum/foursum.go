package foursum

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sum := make([][]int, 0)
	sort.Ints(nums)

	n := len(nums)

	for i := 0; i < n - 3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n - 2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				if l > j+1 && nums[l] == nums[l-1] {
					l++
					continue
				}
				if r < n-1 && nums[r] == nums[r+1] {
					r--
					continue
				}
				s := nums[i] + nums[j] + nums[l] + nums[r]
				if s == target {
					sum = append(sum, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				} else if s < target {
					l++
				} else {
					r--
				}
			}

		}
	}
	return sum
}
