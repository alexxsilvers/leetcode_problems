package main

import "log"

// Given the root node of a binary search tree, return the sum of values of all
// nodes with value between L and R (inclusive).
//
// The binary search tree is guaranteed to have unique values.
func main() {
	ex1 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}
	log.Println(rangeSumBST(ex1, 7, 15)) // 32
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	return dfs(root, L, R)
}

func dfs(root *TreeNode, L, R int) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left, L, R)
	right := dfs(root.Right, L, R)

	s := 0
	if inBorder(root.Val, L, R) {
		s += root.Val
	}

	return s + left + right
}

func inBorder(a, L, R int) bool {
	if a >= L && a <= R {
		return true
	}
	return false
}
