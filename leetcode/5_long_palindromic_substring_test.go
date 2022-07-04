package leetcode

import "testing"

func TestLongPalindromicSubstring(t *testing.T) {
	input := "aacabdkacaa"
	output := "aca"
	t.Logf("input %v, expected output %v", input, output)
	res := LongPalindromicSubstring(input)
	if res == output {
		t.Logf("actual output %v", res)
	} else {
		t.Errorf("actual output %v", res)
	}
}
