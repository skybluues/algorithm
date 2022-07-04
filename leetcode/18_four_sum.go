package leetcode

import "sort"

func FourSum(nums []int, target int) [][]int {
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for i > 0 && i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
		for j := i + 1; j < len(nums); j++ {
			for j > i+1 && j < len(nums) && nums[j] == nums[j-1] {
				j++
			}
			k := j + 1
			h := len(nums) - 1
			for k < len(nums) && k < h {
				sum := nums[i] + nums[j] + nums[k] + nums[h]
				if sum > target {
					h--
				} else if sum < target {
					k++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[h]})
					h--
					k++
				}
				for k > j+1 && k < h && nums[k] == nums[k-1] {
					k++
				}
				for h < len(nums)-1 && h > k && nums[h] == nums[h+1] {
					h--
				}
			}
		}
	}
	return res
}
