package leetcode

func SwapNodesInPairs(head *ListNode) *ListNode {
	res := new(ListNode)
	p := res
	for head != nil && head.Next != nil {
		next := head.Next
		nnext := next.Next
		p.Next = next
		next.Next = head
		head.Next = nnext
		p = head
		head = nnext
	}
	p.Next = head
	return res.Next
}
