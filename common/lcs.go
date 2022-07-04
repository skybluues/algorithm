package algorithm

func LCS(s1 string, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}
	var mat [][]int
	for i := 0; i < len(s1)+1; i++ {
		row := make([]int, len(s2)+1)
		mat = append(mat, row)
	}
	lcs_loc := -1
	lcs_max_len := 0
	for i, r := range s1 {
		for j, c := range s2 {
			if r == c {
				mat[i+1][j+1] = mat[i][j] + 1
				if mat[i+1][j+1] > lcs_max_len {
					lcs_max_len = mat[i+1][j+1]
					lcs_loc = i
				}
			}
		}
	}
	if lcs_max_len > 0 {
		start := lcs_loc - lcs_max_len + 1
		end := lcs_loc + 1
		return s1[start:end]
	} else {
		return ""
	}
}
