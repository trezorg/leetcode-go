package ranges

const maxValue = 10*10*10*10*10 + 1
const minValue = -10*10*10*10*10 - 1

func nextMinFromArrays(nums *[][]int, positions *[]int) (int, int) {

	resArr := -1
	min := maxValue

	for i := 0; i < len(*nums); i++ {
		arr := (*nums)[i]
		arrPosition := (*positions)[i]
		if arrPosition >= len(arr) {
			continue
		}
		value := arr[arrPosition]
		if value < min {
			resArr = i
			min = value
		}
	}

	if resArr != -1 {
		(*positions)[resArr]++
	}

	return min, resArr

}

type valuesList struct {
	values []int
	max    int
	min    int
	full   bool
	added  int
}

func minMax(values []int) (int, int) {
	min := maxValue
	max := minValue
	for i := 0; i < len(values); i++ {
		v := values[i]
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return min, max
}

func (v *valuesList) add(pos int, val int) {
	prevVal := v.values[pos]
	if !v.full && prevVal == minValue {
		v.added++
		if v.added == len(v.values) {
			v.full = true
		}
	}
	v.values[pos] = val
	if v.min == prevVal || v.max == prevVal {
		v.min, v.max = minMax(v.values)
	}
}

func newValuesList(k int) *valuesList {
	values := make([]int, k, k)
	for i := 0; i < k; i++ {
		values[i] = minValue
	}
	return &valuesList{
		min:    maxValue,
		max:    minValue,
		values: values,
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func checkRange(max int, min int, ml int, mr int) bool {
	ln1, ln2 := abs(max-min), abs(mr-ml)
	if ln1 > ln2 {
		return false
	}
	if ln1 < ln2 {
		return true
	}
	return min < ml
}

func smallestRange(nums [][]int) []int {
	k := len(nums)
	positions := make([]int, k, k)
	lst := newValuesList(k)
	ml, mr := minValue, maxValue
	val, arr := nextMinFromArrays(&nums, &positions)
	for arr != -1 {
		lst.add(arr, val)
		if lst.full && checkRange(lst.max, lst.min, ml, mr) {
			ml, mr = lst.min, lst.max
		}
		val, arr = nextMinFromArrays(&nums, &positions)
	}
	if ml == minValue {
		return []int{}
	}
	return []int{ml, mr}
}
