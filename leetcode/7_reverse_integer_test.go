package leetcode

import "testing"

func TestReverseInteger(t *testing.T) {
	x := -123
	expected := -321
	t.Logf("input: %v, expected: %v", x, expected)
	res := ReverseInteger(x)
	if res == expected {
		t.Logf("successful, output: %v", res)
	} else {
		t.Errorf("failed, output: %v", res)
	}
}

func TestReverseInteger2(t *testing.T) {
	x := 120
	expected := 21
	t.Logf("input: %v, expected: %v", x, expected)
	res := ReverseInteger(x)
	if res == expected {
		t.Logf("successful, output: %v", res)
	} else {
		t.Errorf("failed, output: %v", res)
	}
}

func TestReverseInteger3(t *testing.T) {
	x := 12
	expected := 21
	t.Logf("input: %v, expected: %v", x, expected)
	res := ReverseInteger(x)
	if res == expected {
		t.Logf("successful, output: %v", res)
	} else {
		t.Errorf("failed, output: %v", res)
	}
}
