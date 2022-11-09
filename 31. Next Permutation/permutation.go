package main

import "sort"

func nextPermutation(nums []int) {
	k := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			k = i
		}
	}
	if k == -1 {
		sort.Ints(nums)
		return
	}

	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[k] < nums[i] {
			l = i
		}
	}

	nums[k], nums[l] = nums[l], nums[k]

	for i, j := k+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

}
