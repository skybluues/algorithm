package leetcode

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < len(nums) && k > j {
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
				for k > j && nums[j] == nums[j-1] {
					j++
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for k > j && nums[j] == nums[j-1] {
					j++
				}
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return res
}
