package leetcode

import "testing"

func TestMergeTwoSortedLists(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 2
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 4
	l2 := new(ListNode)
	l2.Val = 1
	l2.Next = new(ListNode)
	l2.Next.Val = 3
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4

	res := MergeTwoSortedLists(l1, l2)
	PrintList(l1)
	PrintList(l2)
	PrintList(res)
}
