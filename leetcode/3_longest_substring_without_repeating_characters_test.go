package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	n := LengthOfLongestSubstring(s)
	if n == 3 {
		t.Logf("longth of longest substring of the string \"abcabcbb\" is %d", n)
	} else {
		t.Errorf("input string \"abcabcbb\", expected output %d, real output %d", 3, n)
	}
}
