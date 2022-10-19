package threesum

import (
	"sort"
)

func twoSum(numbers []int, value int) [][]int {
	var sum [][]int
	last := len(numbers) - 1
	for i, j := 0, last; i < j; {
		if numbers[i] > value {
			break
		}
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
			sum = append(sum, []int{numbers[i], numbers[j]})
			i++
			j--
		} else if s > value {
			j--
		} else {
			i++
		}
	}
	return sum
}

func threeSum(numbers []int) [][]int {
	sum := make([][]int, 0)
	sort.Ints(numbers)

	for i := 0; i < len(numbers)-2; i++ {
		if numbers[i] > 0 {
			break
		}
		if i > 0 && numbers[i] == numbers[i-1] {
			continue
		}
		for _, t := range twoSum(numbers[i+1:], 0-numbers[i]) {
			sum = append(sum, []int{numbers[i], t[0], t[1]})
		}
	}
	return sum
}
