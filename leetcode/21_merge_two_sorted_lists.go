package leetcode

func MergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	h := new(ListNode)
	cur := h
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			cur = cur.Next
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = cur.Next
			list2 = list2.Next
		}
	}
	if list1 == nil {
		list1 = list2
	}
	cur.Next = list1
	return h.Next
}
