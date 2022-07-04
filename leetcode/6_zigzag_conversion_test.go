package leetcode

import "testing"

// input s: "PAYPALISHIRING", input numRows: 3
// output:
//   P A H N
//   APLSIIG
//   Y I R

func TestZigzagConversion(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	output := "PAHNAPLSIIGYIR"
	t.Logf("input string: %v", s)
	t.Logf("input numRows: %v", numRows)
	t.Logf("expected output: %v", output)
	res := ZigzagConversion(s, numRows)
	t.Logf("real output: %v", res)
	if res == output {
		t.Logf("result: successful")
	} else {
		t.Errorf("result: failed")
	}
}

func TestZigzagConversion2(t *testing.T) {
	s := "AB"
	numRows := 1
	output := "AB"
	t.Logf("input string: %v", s)
	t.Logf("input numRows: %v", numRows)
	t.Logf("expected output: %v", output)
	res := ZigzagConversion(s, numRows)
	t.Logf("real output: %v", res)
	if res == output {
		t.Logf("result: successful")
	} else {
		t.Errorf("result: failed")
	}
}
