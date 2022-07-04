package leetcode

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	src := x
	const intMax = 2147483647
	res := 0
	for x > 0 {
		digital := x % 10
		x /= 10
		if res > intMax/10 || res*10 > (intMax-digital) {
			return false
		}
		res = res*10 + digital
	}
	if res == src {
		return true
	}
	return false
}
