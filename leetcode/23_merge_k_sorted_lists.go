package leetcode

import "container/heap"

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MergeKSortedLists(lists []*ListNode) *ListNode {
	h := NodeHeap{}
	res := new(ListNode)
	p := res
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		h = append(h, lists[i])
		lists[i] = lists[i].Next
	}
	if len(h) == 0 {
		return nil
	}
	heap.Init(&h)
	for len(h) > 0 {
		t := heap.Pop(&h)
		p.Next = t.(*ListNode)
		p = p.Next
		if p.Next != nil {
			heap.Push(&h, p.Next)
		}
	}
	return res.Next
}
