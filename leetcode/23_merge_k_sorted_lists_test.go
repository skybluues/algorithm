package leetcode

import "testing"

func TestMergeKSortedLists(t *testing.T) {
	l1 := MakeList([]int{1, 4, 5})
	l2 := MakeList([]int{1, 3, 4})
	l3 := MakeList([]int{2, 6})
	input := []*ListNode{l1, l2, l3}
	output := MergeKSortedLists(input)

	PrintList(l1)
	PrintList(l2)
	PrintList(l3)
	PrintList(output)
}
