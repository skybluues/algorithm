package leetcode

import (
	"sort"
)

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func ThreeSumCloset(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1
		for j < len(nums) && j < k {
			val := nums[i] + nums[j] + nums[k]
			if abs(val-target) < abs(res-target) {
				res = val
			}
			if val-target > 0 {
				k--
			} else if val-target < 0 {
				j++
			} else {
				return res
			}
		}
	}
	return res
}
