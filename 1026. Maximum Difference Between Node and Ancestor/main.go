package main

import "log"

// Given the root of a binary tree, find the maximum value V for which there exists
// different nodes A and B where V = |A.val - B.val| and A is an ancestor of B.
//
// (A node A is an ancestor of B if either: any child of A is equal to B,
// or any child of A is an ancestor of B.)
func main() {
	ex1 := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		Right: &TreeNode{
			Val: 10,
			Right: &TreeNode{
				Val: 14,
				Right: &TreeNode{
					Val: 13,
				},
			},
		},
	}
	log.Println(maxAncestorDiff(ex1)) // 7
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}
func dfs(root *TreeNode, mx, mn int) int {
	if root == nil {
		return mx - mn
	}
	mx = max(root.Val, mx)
	mn = min(root.Val, mn)
	return max(dfs(root.Left, mx, mn), dfs(root.Right, mx, mn))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
