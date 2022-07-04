package leetcode

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	t.Logf("nums1 %v", nums1)
	t.Logf("nums2 %v", nums2)
	res := FindMedianSortedArrays(nums1, nums2)
	t.Logf("output is %v", res)
	if res != 2.5 {
		t.Errorf("expected res output is %v", 2.5)
	}
}
