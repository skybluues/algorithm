package leetcode

func RemoveElement(nums []int, val int) int {
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] == val {
			j := i + 1
			for ; j < len(nums); j++ {
				if nums[j] != val {
					break
				}
			}
			if j != len(nums) {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				break
			}
		}
	}
	return i
}
