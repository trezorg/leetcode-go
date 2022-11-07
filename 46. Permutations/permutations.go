package main

func factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

func permute(nums []int) [][]int {

	res := make([][]int, 0, factorial(len(nums)))

	var gen func(a []int, size int)

	gen = func(a []int, size int) {

		if size == 1 {
			c := make([]int, len(a))
			copy(c, a)
			res = append(res, c)
			return
		}

		for i := 0; i < size; i++ {
			gen(a, size-1)
			if size%2 == 1 {
				a[0], a[size-1] = a[size-1], a[0]
			} else {
				a[i], a[size-1] = a[size-1], a[i]
			}
		}
	}

	gen(nums, len(nums))

	return res
}

func permuteIterable(nums []int) [][]int {

	n := len(nums)
	r := len(nums)
	res := make([][]int, 0, factorial(n)/factorial(n-r))

	indices := make([]int, n)
	cycles := make([]int, r)

	for i := 0; i < n; i++ {
		indices[i] = i
	}
	for i := n; i > n-r; i-- {
		cycles[n-i] = i
	}

	values := func() []int {
		values := make([]int, r)
		for i := 0; i < r; i++ {
			values[i] = nums[indices[i]]
		}
		return values
	}

	res = append(res, values())

	for {
		exit := true
		for i := r - 1; i >= 0; i-- {
			cycles[i] -= 1
			if cycles[i] == 0 {
				temp := indices[i]
				for t := i; t < n-1; t++ {
					indices[t] = indices[t+1]
				}
				indices[n-1] = temp
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]
				res = append(res, values())
				exit = false
				break
			}
		}
		if exit {
			break
		}
	}
	return res

}
