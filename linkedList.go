package main

import (
	"fmt"
	//"time"
)

type Node struct {
	Element int
	Next    *Node
}
type LinkedList struct {
	Head   *Node
	Length int
}

func NewLinkedList() *LinkedList {
	list := &LinkedList{
		Head: &Node{
			Element: 1,
		},
		Length: 1,
	}
	return list
}
func (list *LinkedList) InsertHead(node *Node) bool {
	if list.Head == nil {
		fmt.Println("list is empty")
		list.Head = node
		list.Length = 1
		return true
	}
	current := list.Head
	fmt.Println("head:", list.Head.Element, " current:", current.Element)
	node.Next = current
	fmt.Println(" current:", current.Element, " node.next:", node.Next.Element)

	list.Head = node
	fmt.Println("head:", list.Head.Element, " current:", current.Element, "next:", list.Head.Next.Element)
	list.Length += 1
	return true
}
func (list *LinkedList) Travers() {
	if list.Head == nil {
		fmt.Println("List is empty")
		return
	}
	current := list.Head
	for current != nil {
		fmt.Println(" ", current.Element, " ")
		current = current.Next
	}

}
func (list *LinkedList) ReversTravers() {
	pre := list.Head
	current := pre.Next
	var tmp *Node
	for current != nil {
		tmp = current.Next
		current.Next = pre
		pre = current
		current = tmp

	}
	list.Head.Next = nil
	list.Head = pre

}
func (list *LinkedList) DetectLoop() bool {
	slow := list.Head
	fast := list.Head
	for fast != nil && fast != slow {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
	}
	if fast == nil {
		return false
	}
	fast = list.Head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	fmt.Println("Detect the loop entry:", fast.Element)
	return true
}

//func main() {
//	list := NewLinkedList()
//	list.Travers()
//	node := &Node{Element: 2}
//	list.InsertHead(node)
//	list.InsertHead(&Node{Element: 3})

//	list.Travers()
//	list.ReversTravers()
//	list.Travers()

//}
