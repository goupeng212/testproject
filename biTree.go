package main

import (
	"fmt"
	//"time"
)

type BiTreeNode struct {
	Left  *BiTreeNode
	Data  int
	Right *BiTreeNode
}
type BiTree struct {
	Root *BiTreeNode
}

func NewBiTreeNode(data int) *BiTreeNode {
	return &BiTreeNode{Data: data}
}
func NewBiTree() *BiTree {
	return &BiTree{}
}

func (tree *BiTree) Insert(node *BiTreeNode) {
	if tree.Root == nil {
		tree.Root = node
		return
	}
	tree.Root.Insert(node)
}

func (node *BiTreeNode) Insert(v *BiTreeNode) {
	if v.Data < node.Data {
		if node.Left != nil {
			node.Left.Insert(v)
		} else {
			node.Left = v
		}
	} else {
		if node.Right != nil {
			node.Right.Insert(v)
		} else {
			node.Right = v
		}
	}
}

func (tree *BiTree) InOrder() {
	fmt.Print("InOder Travers BiTree: ")
	tree.Root.InOrder()
}

func (node *BiTreeNode) InOrder() {
	if node.Left != nil {
		node.Left.InOrder()
	}
	fmt.Print(node.Data, ",")
	if node.Right != nil {
		node.Right.InOrder()
	}
}
func (tree *BiTree) PreOrder() {
	fmt.Print("PreOder Travers BiTree: ")
	tree.Root.PreOrder()
}

func (node *BiTreeNode) PreOrder() {
	fmt.Print(node.Data, ",")
	if node.Left != nil {
		node.Left.InOrder()
	}

	if node.Right != nil {
		node.Right.InOrder()
	}

}

func CreateBiTree(arr []int) *BiTree {
	biT := NewBiTree()
	for _, ar := range arr {
		node := &BiTreeNode{Data: ar}
		biT.Insert(node)
	}
	return biT

}

func breathTraverse(root *BiTreeNode) {
	var nodes []*BiTreeNode
	if root == nil {
		return
	}
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		fmt.Println(node.Data)
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}

}

//func main() {
//	array := []int{10, 3, 5, 6, 1, 12, 9, 18, 4, 7}
//	biTree := CreateBiTree(array)
//	biTree.InOrder()
//	biTree.PreOrder()
//	fmt.Println("Root:", biTree.Root.Data)

//}
