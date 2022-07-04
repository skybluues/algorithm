package leetcode

func findKth(nums1 []int, nums2 []int, k int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 > len2 {
		return findKth(nums2, nums1, k)
	}
	if len1 == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		if nums1[0] > nums2[0] {
			return nums2[0]
		} else {
			return nums1[0]
		}
	}
	k1 := 0
	k2 := 0
	if k/2 > len1 {
		k1 = len1
	} else {
		k1 = k / 2
	}
	k2 = k - k1
	if nums1[k1-1] > nums2[k2-1] {
		return findKth(nums1, nums2[k2:], k-k2)
	} else if nums1[k1-1] < nums2[k2-1] {
		return findKth(nums1[k1:], nums2, k-k1)
	} else {
		return nums1[k1-1]
	}
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	l := len1 + len2
	if l%2 == 0 {
		return (float64(findKth(nums1, nums2, l/2)) + (float64(findKth(nums1, nums2, l/2+1)))) / 2.0
	} else {
		return float64(findKth(nums1, nums2, l/2+1))
	}
}
