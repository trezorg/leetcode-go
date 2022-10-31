package threesum

import (
	"sort"
)

func twoSum(numbers []int, value int) int {
	var sum int
	diff := int(^uint(0) >> 1)
	last := len(numbers) - 1
	for i, j := 0, last; i < j; {
		if i > 0 && numbers[i] == numbers[i-1] {
			i++
			continue
		}
		if j < last && numbers[j] == numbers[j+1] {
			j--
			continue
		}
		s := numbers[i] + numbers[j]
		if s == value {
			return s
		}
		cDiff := abs(value - s)
		if cDiff < diff {
			diff = cDiff
			sum = s
		}
		if s > value {
			j--
		} else {
			i++
		}
	}
	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func summa(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	res := a[0]
	for i := 1; i < len(a); i++ {
		res += a[i]
	}
	return res
}

func threeSumClosest(nums []int, target int) int {

	diff := int(^uint(0) >> 1)
	var sum int
	sort.Ints(nums)

	if len(nums) >= 3 {
		firstSum, lastSum := summa(nums[:3]...), summa(nums[len(nums)-3:]...)
		if len(nums) == 3 {
			return firstSum
		}
		if firstSum >= target {
			return firstSum
		}
		if lastSum <= target {
			return lastSum
		}
	}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		t := twoSum(nums[i+1:], target-nums[i])
		cDiff := abs(target - (t + nums[i]))
		if cDiff < diff {
			diff = cDiff
			sum = t + nums[i]
		}
	}
	return sum
}
