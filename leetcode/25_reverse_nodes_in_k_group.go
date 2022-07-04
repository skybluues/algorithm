package leetcode

func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		next := head.Next
		head.Next = res
		res = head
		head = next
	}
	return res
}

func ReverseNodesInKGroup(head *ListNode, k int) *ListNode {
	res := new(ListNode)
	p := res
	next := head
	var front *ListNode
	var back *ListNode

	for next != nil {
		front = next
		back = next
		for i := 1; i < k; i++ {
			back = back.Next
			if back == nil {
				break
			}
		}
		if back == nil {
			p.Next = front
			break
		} else {
			next = back.Next
			back.Next = nil
			x := reverseList(front)
			p.Next = x
			p = front
		}
	}
	return res.Next
}
