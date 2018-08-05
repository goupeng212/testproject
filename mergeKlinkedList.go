package main

import (
	"container/heap"
	//"fmt"
	"math"
	//"time"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head, tail *ListNode

	minHeap := &IntHeap{math.MinInt32}
	heap.Init(minHeap)
	for i := 0; i < len(lists); i++ {
		hlist := lists[i]
		for hlist != nil {
			heap.Push(minHeap, hlist.Val)
			hlist = hlist.Next

		}

	}
	//	fmt.Println("intheap:", minHeap.Len())

	if minHeap.Len() > 0 {

		v := heap.Pop(minHeap).(int)
		node := &ListNode{Val: v}
		head = node
		tail = node

	}

	for minHeap.Len() > 0 {

		v := heap.Pop(minHeap).(int)
		node := &ListNode{Val: v}

		tail.Next = node
		tail = node
		//		fmt.Println("intheap:", minHeap.Len())

		//		time.Sleep(time.Second * 2)

	}

	return head.Next
}

//func main() {

//	var lists []*ListNode

//	head1 := &ListNode{Val: 4}
//	node := &ListNode{Val: 3}
//	head1.Next = node
//	node.Next = nil

//	lists = append(lists, head1)

//	head2 := &ListNode{Val: 6}
//	node = &ListNode{Val: 2}
//	head2.Next = node
//	node.Next = nil

//	lists = append(lists, head2)

//	head3 := &ListNode{Val: 1}
//	node = &ListNode{Val: 4}
//	head3.Next = node
//	node.Next = nil

//	lists = append(lists, head3)
//	head := mergeKLists(lists)
//	for head != nil {
//		fmt.Print(head.Val, " ")
//		head = head.Next
//	}
//}
