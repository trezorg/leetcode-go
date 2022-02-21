package oddevenlinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromList(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	result := FromList(data).List()
	assert.Equal(t, data, result)
}

func TestOddEvenList1(t *testing.T) {
	data := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{5, nil}}}}}
	result := oddEvenList(data).List()
	assert.Equal(t, []int{1, 3, 5, 2, 4}, result)
}

func TestOddEvenList2(t *testing.T) {
	data := FromList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	result := oddEvenList(data).List()
	expected := []int{1, 3, 5, 7, 9, 11, 13, 2, 4, 6, 8, 10, 12}
	assert.Equal(t, expected, result)
}

func TestOddEvenList3(t *testing.T) {
	data := FromList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	result := oddEvenList(data).List()
	expected := []int{1, 3, 5, 7, 9, 11, 2, 4, 6, 8, 10, 12}
	assert.Equal(t, expected, result)
}
func TestOddEvenListTwoValue(t *testing.T) {
	data := FromList([]int{1, 2})
	result := oddEvenList(data).List()
	expected := []int{1, 2}
	assert.Equal(t, expected, result)
}
func TestOddEvenListOneValue(t *testing.T) {
	data := FromList([]int{1})
	result := oddEvenList(data).List()
	expected := []int{1}
	assert.Equal(t, expected, result)
}

func TestOddEvenListNilValue(t *testing.T) {
	result := oddEvenList(nil).List()
	assert.Nil(t, result)
}
