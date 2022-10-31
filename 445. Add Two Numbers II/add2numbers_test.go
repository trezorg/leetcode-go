package add2numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateList(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5}
	assert.Equal(t, lst, newList(lst).toList())
}

func TestReverseList(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5}
	list := newList(lst)
	assert.Equal(t, lst, reverse(reverse(list)).toList())
	assert.Equal(t, []int{5, 4, 3, 2, 1}, reverse(list).toList())
}

func TestListToInt(t *testing.T) {
	lst1 := []int{2, 4, 3}
	lst2 := []int{5, 6, 4}
	var lst3 []int
	list1 := newList(lst1)
	list2 := newList(lst2)
	list3 := newList(lst3)
	assert.Equal(t, 243, list1.toInt())
	assert.Equal(t, 564, list2.toInt())
	assert.Equal(t, 0, list3.toInt())
}

func TestAdd2Numbers1(t *testing.T) {
	lst1 := []int{2, 4, 3}
	lst2 := []int{5, 6, 4}
	var lst3 []int
	list1 := newList(lst1)
	list2 := newList(lst2)
	list3 := newList(lst3)
	res1 := addTwoNumbers(list1, list3)
	list2 = newList(lst2)
	list3 = newList(lst3)
	res2 := addTwoNumbers(list2, list3)
	list1 = newList(lst1)
	list2 = newList(lst2)
	res3 := addTwoNumbers(list1, list2)
	assert.Equal(t, []int{2, 4, 3}, res1.toList())
	assert.Equal(t, 243, res1.toInt())
	assert.Equal(t, []int{5, 6, 4}, res2.toList())
	assert.Equal(t, 564, res2.toInt())
	assert.Equal(t, []int{8, 0, 7}, res3.toList())
	assert.Equal(t, 807, res3.toInt())
}

func TestAdd2Numbers2(t *testing.T) {
	lst1 := []int{2, 4, 3, 5}
	lst2 := []int{5, 6, 4}
	list1 := newList(lst1)
	list2 := newList(lst2)
	res := addTwoNumbers(list1, list2)
	assert.Equal(t, []int{2, 9, 9, 9}, res.toList())
	assert.Equal(t, 2999, res.toInt())
}

func TestAdd2Numbers3(t *testing.T) {
	lst1 := []int{5}
	lst2 := []int{5}
	list1 := newList(lst1)
	list2 := newList(lst2)
	res := addTwoNumbers(list1, list2)
	assert.Equal(t, []int{1, 0}, res.toList())
	assert.Equal(t, 10, res.toInt())
	lst1 = []int{7, 7}
	lst2 = []int{7, 7}
	list1 = newList(lst1)
	list2 = newList(lst2)
	res = addTwoNumbers(list1, list2)
	assert.Equal(t, []int{1, 5, 4}, res.toList())
	assert.Equal(t, 154, res.toInt())
}

func TestAdd2Numbers4(t *testing.T) {
	lst1 := []int{1, 8}
	lst2 := []int{0}
	list1 := newList(lst1)
	list2 := newList(lst2)
	res := addTwoNumbers(list1, list2)
	assert.Equal(t, []int{1, 8}, res.toList())
	assert.Equal(t, 18, res.toInt())
}
