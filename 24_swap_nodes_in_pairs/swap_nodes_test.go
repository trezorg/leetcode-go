package swapnodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapNodesEvenNumberOfNodes(t *testing.T) {

	data := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	newData := swapPairs(data)
	expectedList := []int{2, 1, 4, 3, 6, 5}
	assert.Equal(t, expectedList, newData.List())

}

func TestSwapNodesOddNumberOfNodes(t *testing.T) {

	data := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	newData := swapPairs(data)
	expectedList := []int{2, 1, 4, 3, 5}
	assert.Equal(t, expectedList, newData.List())

}

func TestSwapNodesOneElement(t *testing.T) {

	data := &ListNode{Val: 1, Next: nil}
	newData := swapPairs(data)
	expectedList := []int{1}
	assert.Equal(t, expectedList, newData.List())

}

func TestSwapNodesNoElements(t *testing.T) {

	data := (*ListNode)(nil)
	newData := swapPairs(data)
	expectedList := []int{}
	assert.Equal(t, expectedList, newData.List())

}
