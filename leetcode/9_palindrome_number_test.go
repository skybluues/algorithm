package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	input := 121
	expected := true
	output := IsPalindrome(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}

func TestIsPalindrome2(t *testing.T) {
	input := -121
	expected := false
	output := IsPalindrome(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}

func TestIsPalindrome3(t *testing.T) {
	input := 123
	expected := false
	output := IsPalindrome(input)
	t.Logf("input: %v", input)
	t.Logf("expected: %v", expected)
	t.Logf("output: %v", output)
	if expected == output {
		t.Logf("successful")
	} else {
		t.Errorf("failed")
	}
}
