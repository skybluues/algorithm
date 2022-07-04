package leetcode

func RemoveNthNodeFromEndOfList(head *ListNode, n int) *ListNode {
	left := new(ListNode)
	right := left
	left.Next = head
	t := left
	for n > 0 && right != nil {
		n--
		right = right.Next
	}
	if right == nil {
		return head
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return t.Next
}
