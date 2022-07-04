package leetcode

import "testing"

func TestLetterCombinationsOfAPhoneNumber(t *testing.T) {
	input := "23"
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	output := LetterCombinationsOfAPhoneNumber(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
}
