package leetcode

func IntegerToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	str := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			res += str[i]
		}
	}
	return res
}
