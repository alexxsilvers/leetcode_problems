package main

import "github.com/davecgh/go-spew/spew"

// Given the root node of a binary search tree (BST) and a value to be inserted into the tree,
// insert the value into the BST. Return the root node of the BST after the insertion. It is
// guaranteed that the new value does not exist in the original BST.
//
// Note that there may exist multiple valid ways for the insertion, as long as the tree
// remains a BST after insertion. You can return any of them.
func main() {
	ex1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}

	spew.Dump(insertIntoBST(ex1, 5))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
