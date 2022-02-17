package swapnodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) List() []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func find(head *ListNode, k int) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	prev := (*ListNode)(nil)
	i := 0
	for head != nil && i < k-1 {
		prev, head = head, head.Next
		i++
	}
	if i+1 != k {
		return nil, nil
	}
	return prev, head
}

func len(head *ListNode) int {
	i := 0
	for head != nil {
		i++
		head = head.Next
	}
	return i
}

func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil || k < 1 || head.Next == nil {
		return head
	}
	length := len(head)
	median := length/2 + length%2
	if k > median {
		k = length - k + 1
	}
	if k > length || (k == 1 && k == length) {
		return head
	}
	parent_forward, current_forward := find(head, k)
	parent_backward, current_backward := find(head, length-k+1)
	if current_forward == current_backward {
		return head
	}

	is_border_choice := k == 1 || k == length

	if !is_border_choice && current_forward.Next != current_backward {
		current_forward.Next, current_backward.Next = current_backward.Next, current_forward.Next
		parent_forward.Next, parent_backward.Next = current_backward, current_forward
		return head
	}

	if !is_border_choice && current_forward.Next == current_backward {
		current_forward.Next, current_backward.Next = current_backward.Next, current_forward
		parent_forward.Next = current_backward
		return head
	}

	if is_border_choice && current_forward.Next != current_backward {
		current_forward.Next, current_backward.Next = current_backward.Next, current_forward.Next
		parent_backward.Next = current_forward
		return current_backward
	}

	if is_border_choice && current_forward.Next == current_backward {
		current_forward.Next, current_backward.Next = current_backward.Next, current_forward
		return current_backward
	}

	return head
}
