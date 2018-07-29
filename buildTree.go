package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) || len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := Insert(preorder, inorder)
	return root

}
func Insert(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) || len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootV := preorder[0]
	root := &TreeNode{Val: rootV}
	if len(preorder) == 1 {
		root.Left = nil
		root.Right = nil
		return root
	}
	i := 0
	for i = 0; i < len(inorder); i++ {
		if rootV == inorder[i] {
			break
		}
	}
	if i == len(inorder) {
		return nil
	}

	leftInorder := inorder[0:i]
	rightInorder := inorder[i+1:]
	rightIndex := 1
Loop:
	for ; rightIndex < len(preorder); rightIndex++ {
		for k := 0; k < len(rightInorder); k++ {
			if preorder[rightIndex] == rightInorder[k] {
				break Loop
			}
		}
	}
	fmt.Println("rignthIndex:", rightIndex)

	preorderLeftL := preorder[1:rightIndex]
	preorderRight := preorder[rightIndex:]
	root.Left = Insert(preorderLeftL, leftInorder)
	root.Right = Insert(preorderRight, rightInorder)
	return root
}

//func main() {
//	pre := []int{3, 9, 20, 15, 7}
//	in := []int{9, 3, 15, 20, 7}
//	fmt.Print(pre)
//	buildTree(pre, in)

//}
