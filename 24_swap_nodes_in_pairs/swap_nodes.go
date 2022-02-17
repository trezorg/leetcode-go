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

func findChild(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	if node.Next == nil {
		return node
	}
	return node.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first, second := head, head.Next
	nextPairHead := second.Next
	second.Next = first
	first.Next = findChild(nextPairHead)
	head = second
	if first.Next == nil {
		return head
	}
	first = nextPairHead
	for {
		if first == nil || first.Next == nil {
			break
		}
		second = first.Next
		nextPairHead := second.Next
		second.Next = first
		first.Next = findChild(nextPairHead)
		if first.Next == nil {
			break
		}
		first = nextPairHead
	}

	return head

}
