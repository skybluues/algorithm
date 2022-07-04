package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	input := "hello"
	input2 := "ll"
	expected := 2
	output := StrStr(input, input2)
	t.Logf("input: %v, input2: %v", input, input2)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}
