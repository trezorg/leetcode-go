package oddevenlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) List() []int {
	var res []int
	cur := node
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

func FromList(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}
	node := &ListNode{Val: l[0]}
	cur := node
	for i := 1; i < len(l); i++ {
		cur.Next = &ListNode{Val: l[i]}
		cur = cur.Next
	}
	return node
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd, first_even := head, head.Next
	if first_even == nil {
		return head
	}

	even := first_even

	for odd != nil && even != nil {
		odd.Next = even.Next
		if even.Next == nil {
			break
		}
		odd = even.Next
		even.Next = odd.Next
		even = odd.Next
	}

	odd.Next = first_even

	return head

}
