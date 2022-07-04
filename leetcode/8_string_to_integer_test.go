package leetcode

import (
	"testing"
)

func TestStringToInteger(t *testing.T) {
	input := "   -42"
	expected := -42
	output := StringToInteger(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}
