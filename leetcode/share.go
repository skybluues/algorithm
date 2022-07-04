package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v -> ", head.Val)
		head = head.Next
	}
	fmt.Print("nil\n")
}

func MakeList(nums []int) *ListNode {
	var res *ListNode
	if len(nums) == 0 {
		return res
	}
	p := new(ListNode)
	p.Val = nums[0]
	res = p
	for i := 1; i < len(nums); i++ {
		x := new(ListNode)
		x.Val = nums[i]
		p.Next = x
		p = x
	}
	return res
}
