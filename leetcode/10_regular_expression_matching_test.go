package leetcode

import "testing"

func TestRegularExpressionMatching(t *testing.T) {
	input := "ab"
	input2 := ".*"
	expected := true
	output := RegularExpressionMatching(input, input2)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}

func TestRegularExpressionMatching2(t *testing.T) {
	input := "aa"
	input2 := "a*"
	expected := true
	output := RegularExpressionMatching(input, input2)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}
