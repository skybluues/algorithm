package leetcode

func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	needlelen := len(needle)
	for i := 0; i < len(haystack); i++ {
		end := i + needlelen
		if end > len(haystack) {
			end = len(haystack)
		}
		if haystack[i:end] == needle {
			return i
		}
	}
	return -1
}
