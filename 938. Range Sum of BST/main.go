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

// TreeNode definition of the node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	left := rangeSumBST(root.Left, L, R)
	right := rangeSumBST(root.Right, L, R)

	if inBounds(root.Val, L, R) {
		return root.Val + left + right
	}

	return left + right
}

func inBounds(v, L, R int) bool {
	return v >= L && v <= R
}
