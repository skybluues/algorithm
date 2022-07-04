package leetcode

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := ""
	for i := 0; i < len(strs[0]); i++ {
		t := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i {
				return res
			}
			if t != strs[j][i] {
				return res
			}
		}
		res += string(t)
	}
	return res
}
