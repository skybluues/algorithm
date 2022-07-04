package leetcode

import "testing"

func TestReverseNodesInKGroup(t *testing.T) {
	head := MakeList([]int{1, 2, 3, 4, 5})
	k := 2
	PrintList(head)
	output := ReverseNodesInKGroup(head, k)
	PrintList(output)
}
