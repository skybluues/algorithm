package leetcode

func RemoveDuplicatesFromSortedArray(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	t := 1
	r := 1
	for t < len(nums) {
		if nums[t] > nums[t-1] {
			t++
			continue
		}
		for r = t + 1; r < len(nums); r++ {
			if nums[r] > nums[t-1] {
				break
			}
		}
		if r == len(nums) {
			break
		} else {
			nums[t] = nums[r]
			t++
		}
	}
	return t
}
