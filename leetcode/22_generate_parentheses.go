package leetcode

func generateParenthesesHelper(l int, r int, str string, ans *[]string) {
	if l == 0 && r == 0 {
		*ans = append(*ans, str)
	}
	if l > 0 {
		generateParenthesesHelper(l-1, r, str+"(", ans)
	}
	if l < r {
		generateParenthesesHelper(l, r-1, str+")", ans)
	}
}

func GenerateParentheses(n int) []string {
	ans := new([]string)
	l := n
	r := n
	generateParenthesesHelper(l, r, "", ans)
	return *ans
}
