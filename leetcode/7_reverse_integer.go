package leetcode

func ReverseInteger(x int) int {
	const intMax = 2147483647
	const intMin = -2147483648
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > intMax/10 || rev*10 > intMax-pop {
			return 0
		}
		if rev < intMin/10 || rev*10 < intMin-pop {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
