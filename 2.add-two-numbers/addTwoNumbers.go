package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := ListNode{0, nil}
	current := &head
	var carry int
	first := l1
	second := l2
	for {
		v1, v2 := 0, 0

		if first != nil {
			v1 = first.Val
			first = first.Next

		}

		if second != nil {
			v2 = second.Val
			second = second.Next

		}

		total := carry + v1 + v2
		carry = total / 10
		current.Next = &ListNode{total % 10, nil}
		current = current.Next
		if first == nil && second == nil {
			break
		}

	}

	if carry != 0 {
		current.Next = &ListNode{carry, nil}
	}

	return head.Next
}
