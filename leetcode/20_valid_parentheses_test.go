package leetcode

import "testing"

func TestValidParentheses(t *testing.T) {
	input := "()[]{}"
	expected := true
	output := ValidParentheses(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}
