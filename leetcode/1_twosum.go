package leetcode

func TwoSum(array []int, target int) (int, int) {
	one := -1
	two := -1
	cache := make(map[int]int)
	for index, value := range array {
		if v, ok := cache[target-value]; ok != false {
			one = v
			two = index
			break
		} else {
			cache[value] = index
		}
	}
	return one, two
}
