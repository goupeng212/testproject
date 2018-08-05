package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	tail.Next = head

	pre, ahead := head, head
	count := k % length
	for i := 0; i < count; i++ {
		ahead = ahead.Next
	}
	for ahead != tail {
		ahead = ahead.Next
		pre = pre.Next
	}
	head = pre.Next
	pre.Next = nil
	return head

}
