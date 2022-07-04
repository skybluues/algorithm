package leetcode

import "testing"

func TestGenerateParentheses(t *testing.T) {
	input := 3
	expected := []string{"((()))", "(())()", "(()())", "()(())", "()()()"}
	output := GenerateParentheses(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
}
