package leetcode

import (
	"testing"
)

func TestThreeSumCloset(t *testing.T) {
	input := []int{-1, 2, 1, -4}
	target := 1
	expected := 2
	output := ThreeSumCloset(input, target)
	t.Logf("input: %v, target: %v", input, target)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}
