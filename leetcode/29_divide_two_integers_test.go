package leetcode

import "testing"

func TestDivideTwoIntegers(t *testing.T) {
	input := 10
	input2 := 3
	expected := 3
	output := DivideTwoInteger(input, input2)
	t.Logf("input: %v / %v", input, input2)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}
