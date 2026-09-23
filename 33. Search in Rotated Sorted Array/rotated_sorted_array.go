package main

// returns index of pivot in rotated sorted array
// [4,5,6,7,0,1,2] -> 4
// [5,6,7,0,1,2,4] -> 3
func searchPivot(nums []int) int {
	return searchPivotIndex(nums, 0, len(nums)-1)
}

func searchPivotIndex(nums []int, i, j int) int {
	if i == j {
		return 0
	}
	if j-i == 1 {
		return j
	}
	midIndex := (i + j) / 2
	mid, start, end := nums[midIndex], nums[i], nums[j]
	// in the middle of the sorted slice
	if mid >= start && mid <= end {
		return i
	}
	if mid > end {
		return searchPivotIndex(nums, midIndex, j)
	}
	if mid < start {
		return searchPivotIndex(nums, i, midIndex)
	}
	return 0
}

func binarySearchIndex(nums []int, target, i, j int) int {
	if i == j {
		if nums[i] == target {
			return i
		}
		return -1
	}
	midIndex := (i + j) / 2
	mid := nums[midIndex]
	if mid == target {
		return midIndex
	}
	if mid < target {
		return binarySearchIndex(nums, target, midIndex+1, j)
	}
	if mid > target {
		return binarySearchIndex(nums, target, i, midIndex-1)
	}
	return -1
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	last := len(nums) - 1
	pivot := searchPivotIndex(nums, 0, last)
	if nums[pivot] <= target && nums[last] >= target || pivot == 0 {
		return binarySearchIndex(nums, target, pivot, last)
	}
	return binarySearchIndex(nums, target, 0, pivot-1)
}
