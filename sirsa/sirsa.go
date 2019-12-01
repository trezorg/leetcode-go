package sirsa

func minIdx(nums []int, a int, b int) int {
	if nums[a] < nums[b] {
		return a
	}
	return b
}

func findMin(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return 0
	}

	start := 0
	end := len(nums) - 1

	for start != end {
		if end-start < 2 {
			return minIdx(nums, start, end)
		}
		mid := (end-start)/2 + start
		el := nums[mid]
		if nums[end] < el {
			// find in left
			start = mid
		} else if nums[start] > el {
			// find in right
			end = mid
		} else {
			return start
		}

	}

	return start

}

func findMinPivot(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		if nums[i - 1] > nums[i] {
			return i
		}
	}

	return 0

}

func binSearch(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	start := 0
	end := len(nums) - 1

	for start < end {
		mid := (end-start)/2 + start
		el := nums[mid]
		if el == target {
			return mid
		} else if el < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}

	if nums[start] == target {
		return start
	}

	return -1

}

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}
	minIdx := findMin(nums)
	if target == nums[minIdx] {
		return minIdx
	}

	if nums[len(nums)-1] >= target {
		res := binSearch(nums[minIdx:], target)
		if res == -1 {
			return -1
		}
		return res + minIdx
	}
	return binSearch(nums[:minIdx], target)

}

func search2(nums []int, target int) bool {

	if len(nums) == 0 {
		return false
	}
	minIdx := findMinPivot(nums)
	if target == nums[minIdx] {
		return true
	}

	if nums[len(nums)-1] >= target {
		if binSearch(nums[minIdx:], target) == -1 {
			return false
		}
		return true
	}
	if binSearch(nums[:minIdx], target) == -1 {
		return false
	}
	return true

}
