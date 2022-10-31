package fmp

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] > 0 && nums[i] <= n && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	for i := 0; i < n; {
		if nums[i] != i+1 {
			return i + 1
		}
		i++
	}
	return n + 1

}
