package leetcode

func LongPalindromicSubstring(s string) string {
	if len(s) < 2 {
		return s
	}
	maxlen := 0
	res := ""
	for i := 0; i < len(s); i++ {
		left := i
		right := i
		for {
			if left < 0 || right >= len(s) || s[left] != s[right] {
				break
			}
			left--
			right++
		}
		if right-left-1 > maxlen {
			maxlen = right - left - 1
			res = s[left+1 : right]
		}
		left = i
		right = i + 1
		for {
			if left < 0 || right >= len(s) || s[left] != s[right] {
				break
			}
			left--
			right++
		}
		if right-left-1 > maxlen {
			maxlen = right - left - 1
			res = s[left+1 : right]
		}
	}
	return res
}
