package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func searchRange(nums []int, target int) []int {

	res := []int{len(nums), -1}
	var search func(i, j int)

	search = func(i, j int) {
		if i >= j {
			if i < len(nums) && nums[i] == target {
				res[0] = min(i, res[0])
				res[1] = max(i, res[1])
			}
			return
		}
		middle := (i + j) / 2
		if nums[middle] == target {
			res[0] = min(middle, res[0])
			res[1] = max(middle, res[1])
			search(i, middle-1)
			search(middle+1, j)
		} else if nums[middle] > target {
			search(i, middle-1)
		} else {
			search(middle+1, j)
		}
	}

	search(0, len(nums)-1)
	if res[1] == -1 {
		res[0] = -1
	}
	return res

}
