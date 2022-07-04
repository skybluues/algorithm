package leetcode

import "math"

func DivideTwoInteger(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if divisor == 0 {
		return math.MaxInt
	}
	sign := true
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = false
	}
	x := dividend
	if dividend < 0 {
		x = -dividend
	}
	y := divisor
	if divisor < 0 {
		y = -divisor
	}
	res := 0
	for x >= y {
		shit := 0
		for x >= y*(1<<shit) {
			shit += 1
		}
		res += (1 << (shit - 1))
		x -= y * (1 << (shit - 1))
	}
	if sign {
		return res
	} else {
		return -res
	}
}
