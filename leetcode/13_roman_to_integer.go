package leetcode

func RomanToInteger(s string) int {
	romanDoubleMap := map[string]int{
		"CM": 900,
		"CD": 400,
		"XC": 90,
		"XL": 40,
		"IX": 9,
		"IV": 4,
	}
	romanSingleMap := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if val, ok := romanDoubleMap[s[i:i+2]]; ok {
				res += val
				i++
				continue
			}
		}
		if val, ok := romanSingleMap[s[i:i+1]]; ok {
			res += val
		}
	}
	return res
}
