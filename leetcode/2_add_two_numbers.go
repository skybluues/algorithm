package leetcode

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	h1 := new(ListNode)
	h1.Next = l1
	h2 := new(ListNode)
	h2.Next = l2
	prev1 := h1
	prev2 := h2
	for {
		if l1 != nil && l2 != nil {
			sum := l1.Val + l2.Val + carry
			l1.Val = sum % 10
			carry = sum / 10
			prev1 = prev1.Next
			prev2 = prev2.Next
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 == nil {
			prev1.Next = l2
			l1 = prev1.Next
			break
		} else {
			break
		}
	}
	for {
		if l1 == nil {
			if carry > 0 {
				n := new(ListNode)
				n.Val = carry
				prev1.Next = n
			}
			break
		} else {
			sum := l1.Val + carry
			l1.Val = sum % 10
			carry = sum / 10
			l1 = l1.Next
			prev1 = prev1.Next
		}
	}
	return h1.Next
}
