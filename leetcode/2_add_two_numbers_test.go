package leetcode

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	var l1 *ListNode
	var l2 *ListNode
	for i := 0; i < 7; i++ {
		n := new(ListNode)
		n.Val = 9
		if l1 == nil {
			l1 = n
		} else {
			n.Next = l1
			l1 = n
		}
	}
	for i := 0; i < 4; i++ {
		n := new(ListNode)
		n.Val = 9
		if l2 == nil {
			l2 = n
		} else {
			n.Next = l2
			l2 = n
		}
	}
	res := AddTwoNumbers(l1, l2)
	v := 0
	e := 1
	for ; res != nil; res = res.Next {
		v = v + res.Val*e
		e = e * 10
	}
	if v == 10009998 {
		t.Log(true)
	} else {
		t.Errorf("epected value is 10009998, real value is %d", v)
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 4

	l2 := new(ListNode)
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4

	res := AddTwoNumbers(l1, l2)
	for ; res != nil; res = res.Next {
		t.Log(res.Val)
	}
}
