package leetcode

import "testing"

func TestRomanToInteger(t *testing.T) {
	input := "LVIII"
	expected := 58
	output := RomanToInteger(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}
