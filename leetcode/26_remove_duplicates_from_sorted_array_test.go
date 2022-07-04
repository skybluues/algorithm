package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	input := []int{1, 1}
	output := RemoveDuplicatesFromSortedArray(input)
	fmt.Print(output)
	flag := output == 1 && input[0] == 1
	if flag {
		t.Log("ok")
	} else {
		t.Error("failed")
	}
}
