package leetcode

import "testing"

func TestSwapNodesInPairs(t *testing.T) {
	input := MakeList([]int{1, 2, 3, 4})
	expected := MakeList([]int{2, 1, 4, 3})
	output := SwapNodesInPairs(input)
	PrintList(input)
	PrintList(expected)
	PrintList(output)
}
