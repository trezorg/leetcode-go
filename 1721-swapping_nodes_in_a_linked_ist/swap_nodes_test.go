package swapnodes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapNodes(t *testing.T) {
	data := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	newData := swapNodes(data, 2)
	expectedList := []int{1, 5, 3, 4, 2, 6}
	assert.Equal(t, expectedList, newData.List())
}

func TestSwapNodesOutOfRange(t *testing.T) {
	data := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	newData := swapNodes(data, 10)
	expectedList := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, expectedList, newData.List())
}

func TestSwapNodesFirstElement(t *testing.T) {
	data := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	newData := swapNodes(data, 1)
	expectedList := []int{6, 2, 3, 4, 5, 1}
	assert.Equal(t, expectedList, newData.List())
}

func TestSwapNodesLastElement(t *testing.T) {
	data := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	newData := swapNodes(data, 6)
	expectedList := []int{6, 2, 3, 4, 5, 1}
	assert.Equal(t, expectedList, newData.List())
}

func TestSwapNodesOneElement(t *testing.T) {
	data := &ListNode{Val: 1, Next: nil}
	newData := swapNodes(data, 1)
	expectedList := []int{1}
	assert.Equal(t, expectedList, newData.List())
}

func TestSwapNodesTwoElementFirstChoice(t *testing.T) {
	data := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	newData := swapNodes(data, 1)
	expectedList := []int{2, 1}
	assert.Equal(t, expectedList, newData.List())
}

func TestSwapNodesTwoElementLastChoice(t *testing.T) {
	data := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	newData := swapNodes(data, 2)
	expectedList := []int{2, 1}
	assert.Equal(t, expectedList, newData.List())
}
func TestSwapNodesMiddle(t *testing.T) {
	data := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}}}}}}}
	newData := swapNodes(data, 4)
	expectedList := []int{1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, expectedList, newData.List())
}

func TestSwapNodesSiblings(t *testing.T) {
	data := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	newData := swapNodes(data, 3)
	expectedList := []int{1, 2, 4, 3, 5, 6}
	assert.Equal(t, expectedList, newData.List())
}
