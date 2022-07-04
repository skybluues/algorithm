package algorithm

import (
	"testing"
)

func TestLCS(t *testing.T) {
	s1 := "abcdefggfedcba"
	s2 := "ghtcbaflmyta"
	t.Logf("left hand string %s", s1)
	t.Logf("right hand string %s", s2)
	res := LCS(s1, s2)
	if res == "cba" {
		t.Log("SUCCESSED, common substring cba")
	} else {
		t.Errorf("FAILED, common substring %s", res)
	}
}
