package main

import "log"

// Given a binary tree, find its maximum depth.
//
// The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
// Note: A leaf is a node with no children.
func main() {
	ex1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 7,
			Right: &TreeNode{
				Val: 10,
			},
		},
		Right: nil,
	}

	log.Println(maxDepth(ex1)) // 3
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive solution
// Runtime - O(n)
// Space - O(k) where k is a count of nodes
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}
