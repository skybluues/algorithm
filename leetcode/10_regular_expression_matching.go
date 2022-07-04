package leetcode

func RegularExpressionMatching(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	first_match := len(s) != 0 && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return RegularExpressionMatching(s, p[2:]) || (first_match && RegularExpressionMatching(s[1:], p))
	} else {
		return first_match && RegularExpressionMatching(s[1:], p[1:])
	}
}
