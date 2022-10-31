package rotatelist

import (
	"fmt"
	"strconv"
	"strings"
)

func IntToString(values []int, sep string) string {
	b := make([]string, len(values))
	for i, v := range values {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0], Next: nil}
	last := head
	for i := 1; i < len(values); i++ {
		new := &ListNode{Val: values[i], Next: nil}
		last.Next = new
		last = new
	}
	return head
}

func (head *ListNode) toSlice() []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func (head *ListNode) String() string {
	return fmt.Sprintf("ListNode(Val: %d, Next: %#v)", head.Val, head.Next)
}

func (head *ListNode) toString() string {
	return IntToString(head.toSlice(), "->")
}

func (head *ListNode) tail() *ListNode {
	for head != nil && head.Next != nil {
		head = head.Next
	}
	return head
}

func (head *ListNode) len() int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}

func (head *ListNode) rotate() *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	next := cur.Next
	for next != nil {
		temp := next.Next
		next.Next, cur = cur, next
		next = temp
	}
	head.Next = nil
	return cur
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if k == 0 || head.Next == nil {
		return head
	}

	tail := head
	len := 1
	for tail.Next != nil {
		tail = tail.Next
		len++
	}

	shift := len - (k % len)
	tail.Next = head

	prev := head
	next := head.Next
	for i := 1; i < shift; i++ {
		prev = next
		next = next.Next
	}
	prev.Next = nil
	return next
}
