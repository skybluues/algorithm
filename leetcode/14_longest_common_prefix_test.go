package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	input := []string{"ab", "a"}
	expected := "a"
	output := LongestCommonPrefix(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}
