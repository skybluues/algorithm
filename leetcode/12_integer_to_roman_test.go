package leetcode

import "testing"

func TestIntegerToRoman(t *testing.T) {
	input := 58
	expected := "LVIII"
	output := IntegerToRoman(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}

func TestIntegerToRoman2(t *testing.T) {
	input := 1994
	expected := "MCMXCIV"
	output := IntegerToRoman(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}
