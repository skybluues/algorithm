package main

import (
	"algorithm/leetcode"
	"fmt"
)

func main() {
	nums1 := []int{}
	nums2 := []int{1}
	res := leetcode.FindMedianSortedArrays(nums1, nums2)
	fmt.Printf("res is %v\n", res)
}
