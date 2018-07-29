package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// origin:=head
	if head.Next == nil && n == 1 {
		return nil
	}
	revertHead := revert(head)
	traverLinkedList(revertHead)
	target := revertHead
	pre := revertHead
	for i := 1; i < n; i++ {
		pre = target
		target = target.Next
	}
	if pre == target {
		revertHead = revertHead.Next
	} else {
		pre.Next = target.Next
	}
	head = revert(revertHead)
	return head
}
func revert(head *ListNode) *ListNode {
	pre := head
	current := head.Next
	for current != nil {
		tmp := current.Next
		current.Next = pre
		pre = current
		current = tmp
	}
	head.Next = nil
	return pre
}

func revert2(head *ListNode) *ListNode {
	var pre *ListNode
	current, next := head, head
	for current != nil {
		next = current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}

func InsertNode(head *ListNode, node *ListNode) *ListNode {
	if head == nil {
		//		fmt.Println("head is nil")
		head = node
		node.Next = nil
		//		fmt.Println(head.Val, node.Val)

		return head
	}
	node.Next = head
	head = node
	return head
}
func traverLinkedList(head *ListNode) {
	point := head
	for point != nil {
		fmt.Print(point.Val)
		point = point.Next
	}
}

//func main() {
//	var head *ListNode
//	length := 0
//	arr := []int{1, 3, 2, 1, 3, 15, 7, 4}
//	for _, value := range arr {
//		node := &ListNode{Val: value, Next: nil}
//		head = InsertNode(head, node)
//		//		fmt.Println(head.Val)

//		length++
//	}
//	fmt.Println("length:", length, head.Val)
//	traverLinkedList(head)
//	removeNthFromEnd(head, 1)
//	var leng int
//	leng = 10

//}
