package leetcode

func LetterCombinationsOfAPhoneNumber(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	res := []string{""}
	for i := 0; i < len(digits); i++ {
		tmp := []string{}
		set, _ := m[string(digits[i])]
		for j := 0; j < len(res); j++ {
			t := res[j]
			for k := 0; k < len(set); k++ {
				tmp = append(tmp, t+set[k])
			}
		}
		res = tmp
	}
	return res
}
