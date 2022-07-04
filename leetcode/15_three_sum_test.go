package leetcode

import "testing"

func TestThreeSum(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	output := ThreeSum(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
}
