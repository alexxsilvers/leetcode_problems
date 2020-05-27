package main

import "log"

// Given a non-empty binary search tree and a target value, find the value in the BST that is closest to the target.
// Given target value is a floating point.
// You are guaranteed to have only one unique value in the BST that is closest to the target.
func main() {
	bst := &TreeNode{
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
			Val: 5,
		},
	}

	log.Println(closestValueBinary(bst, 3.7)) // 4
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValueBinary(root *TreeNode, target float64) int {
	v, closest := root.Val, root.Val
	for root != nil {
		v = root.Val
		if abs(float64(v)-target) < abs(float64(closest)-target) {
			closest = v
		}

		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return closest
}

func closestValue(root *TreeNode, target float64) int {
	if root == nil {
		return 0
	}

	closest := root.Val
	diff := abs(float64(root.Val) - target)

	dfs(root, &closest, &diff, target)

	return closest
}

func dfs(node *TreeNode, closest *int, diff *float64, target float64) {
	if node == nil {
		return
	}

	d := abs(float64(node.Val) - target)
	if d < *diff {
		*closest = node.Val
		*diff = d
	}

	dfs(node.Left, closest, diff, target)
	dfs(node.Right, closest, diff, target)
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
