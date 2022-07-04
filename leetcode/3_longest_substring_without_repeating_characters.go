package leetcode

func LengthOfLongestSubstring(s string) int {
	length := 0
	locMap := make(map[byte]int)
	left := 0
	for i := 0; i < len(s); i++ {
		if loc, ok := locMap[s[i]]; ok && loc != -1 {
			thisLen := i - left
			if thisLen > length {
				length = thisLen
			}
			for j := left; j <= loc; j++ {
				locMap[s[j]] = -1
			}
			left = loc + 1
		}
		locMap[s[i]] = i
	}
	thisLen := len(s) - left
	if thisLen > length {
		length = thisLen
	}
	return length
}
