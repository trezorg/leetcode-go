package main

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

func (head *ListNode) Slice() []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	if n > length {
		return head
	}

	position, idx := length-n, 0
	cur = head
	var prev *ListNode = nil
	for cur != nil {
		if idx == position {
			if prev == nil {
				head = head.Next
				break
			}
			prev.Next = cur.Next
			break
		}
		idx++
		prev = cur
		cur = cur.Next
	}
	return head
}
