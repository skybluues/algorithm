package leetcode

func ZigzagConversion(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	zz := [][]byte{}
	for i := 0; i < numRows; i++ {
		zz = append(zz, []byte{})
	}
	curLen := 0
	step := -1
	for i := 0; i < len(s); i++ {
		zz[curLen] = append(zz[curLen], s[i])
		if curLen == 0 || curLen == numRows-1 {
			step = -step
		}
		curLen = curLen + step
	}
	res := ""
	for i := 0; i < numRows; i++ {
		res = res + string(zz[i])
	}
	return res
}
