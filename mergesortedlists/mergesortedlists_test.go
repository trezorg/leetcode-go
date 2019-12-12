package mergesortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists1(t *testing.T) {
	lst1 := []int{2, 3, 4}
	lst2 := []int{3, 4, 5}
	list1 := newList(lst1)
	list2 := newList(lst2)
	res := merge2Lists(list1, list2)
	assert.Equal(t, []int{2, 3, 3, 4, 4, 5}, res.toList())
}

func TestMergeTwoLists2(t *testing.T) {
	lst1 := []int{2, 2, 2}
	lst2 := []int{3, 4, 5}
	list1 := newList(lst1)
	list2 := newList(lst2)
	res := merge2Lists(list1, list2)
	assert.Equal(t, []int{2, 2, 2, 3, 4, 5}, res.toList())
}

func TestMergeTwoLists3(t *testing.T) {
	var lst1 []int
	lst2 := []int{3, 4, 5}
	list1 := newList(lst1)
	list2 := newList(lst2)
	res := merge2Lists(list1, list2)
	assert.Equal(t, []int{3, 4, 5}, res.toList())
}

func TestMergeKLists(t *testing.T) {
	lst1 := []int{2, 3, 4}
	lst2 := []int{3, 4, 5}
	list1 := newList(lst1)
	list2 := newList(lst2)
	res := mergeKLists([]*ListNode{list1, list2})
	assert.Equal(t, []int{2, 3, 3, 4, 4, 5}, res.toList())
}

func TestMergeKLists2(t *testing.T) {
	lst1 := []int{2, 3, 4}
	lst2 := []int{3, 4, 5}
	lst3 := []int{0, 1, 6}
	list1 := newList(lst1)
	list2 := newList(lst2)
	list3 := newList(lst3)
	res := mergeKLists([]*ListNode{list1, list2, list3})
	assert.Equal(t, []int{0, 1, 2, 3, 3, 4, 4, 5, 6}, res.toList())
}
