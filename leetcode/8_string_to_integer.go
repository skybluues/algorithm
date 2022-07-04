package leetcode

func StringToInteger(s string) int {
	l := 0
	for ; l < len(s); l++ {
		if s[l] != byte(' ') {
			break
		}
	}
	if l == len(s) {
		return 0
	}
	flag := 1
	if s[l] == '-' || s[l] == '+' {
		if s[l] == '+' {
			flag = 1
		}
		if s[l] == '-' {
			flag = -1
		}
		l++
	}
	if l == len(s) {
		return 0
	}
	res := 0
	const intMax = 2147483647
	const intMin = -2147483648
	for ; l < len(s); l++ {
		if s[l] < byte('0') || s[l] > byte('9') {
			break
		} else {
			v := s[l] - byte('0')
			if res > (intMax-int(v)*flag)/10 || res*10 > intMax-int(v) {
				return intMax
			} else if res < (intMin-int(v)*flag)/10 || res*10 < intMin-int(v) {
				return intMin
			} else {
				res = res*10 + int(v)*flag
			}
		}
	}
	return res
}
