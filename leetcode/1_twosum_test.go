package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	array := []int{4, 4, 9, 2, 5, 4, 2, 7, 10, 55, 5, 99}
	target := 65
	one, two := TwoSum(array, target)
	if one == 8 && two == 9 {
		t.Logf("SUCCESSED, one is %d, two is %d", one, two)
	} else {
		t.Errorf("FAILED, one is %d, two is %d", one, two)
	}
}
