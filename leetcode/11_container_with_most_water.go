package leetcode

func ContainerWithMostWater(height []int) int {
	size := len(height)
	if size <= 1 {
		return 0
	}
	l := 0
	r := size - 1
	area := 0
	for l < r {
		w := r - l
		h := 0
		if height[l] < height[r] {
			h = height[l]
			l++
		} else {
			h = height[r]
			r--
		}
		a := w * h
		if a > area {
			area = a
		}
	}
	return area
}
