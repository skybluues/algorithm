package leetcode

import "testing"

func TestContainerWithMostWater(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49
	output := ContainerWithMostWater(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}
