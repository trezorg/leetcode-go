package addtwonumbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSingleValues(t *testing.T) {
	data1 := &ListNode{Val: 1, Next: nil}
	data2 := &ListNode{Val: 1, Next: nil}
	result := addTwoNumbers(data1, data2).List()
	assert.Equal(t, []int{2}, result)
}

func TestAddSingleValuesWithZero(t *testing.T) {
	data1 := &ListNode{Val: 1, Next: nil}
	data2 := &ListNode{Val: 0, Next: nil}
	result := addTwoNumbers(data1, data2).List()
	assert.Equal(t, []int{1}, result)
}
func TestAddSingleValuesDifferentSizeOfLists(t *testing.T) {
	data1 := &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: nil}}
	data2 := &ListNode{Val: 7, Next: nil}
	result := addTwoNumbers(data1, data2).List()
	assert.Equal(t, []int{8, 5}, result)
}

func TestAddZeros(t *testing.T) {
	data1 := &ListNode{Val: 0, Next: nil}
	data2 := &ListNode{Val: 0, Next: nil}
	result := addTwoNumbers(data1, data2).List()
	assert.Equal(t, []int{0}, result)
}
func TestAddMediumLists(t *testing.T) {
	data1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	data2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	result := addTwoNumbers(data1, data2).List()
	assert.Equal(t, []int{7, 0, 8}, result)
}
func TestAddLongLists(t *testing.T) {
	data1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}}
	data2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}
	result := addTwoNumbers(data1, data2).List()
	assert.Equal(t, []int{8, 9, 9, 9, 0, 0, 0, 1}, result)
}

func TestAddVeryLists(t *testing.T) {
	data1 := newList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	data2 := newList([]int{5, 6, 4})
	result := addTwoNumbers(data1, data2).List()
	assert.Equal(t, []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, result)
}
