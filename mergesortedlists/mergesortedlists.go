package mergesortedlists

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ListNode is list node
type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) toList() []int {
	var res []int
	cur := node
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

func (node *ListNode) toInt() int {
	var res int
	counter := 0
	cur := node
	for cur != nil {
		res += int(math.Pow10(counter)) * cur.Val
		counter++
		cur = cur.Next
	}
	return res
}

func (node *ListNode) String() string {
	var res []string
	cur := node
	for cur != nil {
		res = append(res, strconv.Itoa(cur.Val))
		cur = cur.Next
	}
	return fmt.Sprintf("(%s)", strings.Join(res, " -> "))
}

func newList(lst []int) *ListNode {
	if len(lst) == 0 {
		return nil
	}
	node := &ListNode{Val: lst[0]}
	res := node
	for i := 1; i < len(lst); i++ {
		newNode := &ListNode{Val: lst[i]}
		node.Next = newNode
		node = newNode
	}
	return res
}

func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var node *ListNode
	if l1.Val < l2.Val {
		node = l1
		l1 = l1.Next
	} else {
		node = l2
		l2 = l2.Next
	}
	root := node

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	} else if l2 != nil {
		node.Next = l2
	}

	return root

}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return merge2Lists(lists[0], lists[1])
	}
	mid := (len(lists) - 1) / 2
	return merge2Lists(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}
