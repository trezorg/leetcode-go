package addtwonumbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	rest := 0

	sm := l1.Val + l2.Val + rest
	n := sm % 10
	rest = sm / 10
	node := &ListNode{Val: n}
	root := node
	l1 = l1.Next
	l2 = l2.Next

	for l1 != nil && l2 != nil {
		sm := l1.Val + l2.Val + rest
		n := sm % 10
		rest = sm / 10
		temp := &ListNode{Val: n}
		node.Next = temp
		node = node.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		if rest == 0 {
			node.Next = l1
			break
		}
		sm := l1.Val + rest
		n := sm % 10
		rest = sm / 10
		temp := &ListNode{Val: n}
		node.Next = temp
		node = node.Next
		l1 = l1.Next
	}
	for l2 != nil {
		if rest == 0 {
			node.Next = l2
			break
		}
		sm := l2.Val + rest
		rest = sm / 10
		n := sm % 10
		temp := &ListNode{Val: n}
		node.Next = temp
		node = node.Next
		l2 = l2.Next
	}

	for rest != 0 {
		n := rest % 10
		rest = rest / 10
		temp := &ListNode{Val: n}
		node.Next = temp
		node = temp
	}

	return root

}
