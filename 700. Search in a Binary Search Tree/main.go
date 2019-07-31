package main

import "github.com/davecgh/go-spew/spew"

// Given the root node of a binary search tree (BST) and a value. You need to
// find the node in the BST that the node's value equals the given value.
// Return the subtree rooted with that node. If such node doesn't exist,
// you should return NULL.
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

	spew.Dump(searchBST(ex1, 5))
	spew.Dump(searchBST(ex1, 2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}

	return nil
}
