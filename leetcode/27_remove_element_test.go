package leetcode

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	input := []int{3, 2, 2, 3}
	val := 3
	output := RemoveElement(input, val)
	fmt.Print(output)
	flag := output == 2 && input[0] == 2 && input[1] == 2
	if flag {
		t.Log("ok")
	} else {
		t.Error("failed")
	}
}
