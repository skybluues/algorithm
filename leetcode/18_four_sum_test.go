package leetcode

import "testing"

func TestFourSum(t *testing.T) {
	input := []int{1, 0, -1, 0, -2, 2}
	target := 0
	expected := [][]int{
		[]int{-2, -1, 1, 2},
		[]int{-2, 0, 0, 2},
		[]int{-1, 0, 0, 1},
	}
	output := FourSum(input, target)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
}
