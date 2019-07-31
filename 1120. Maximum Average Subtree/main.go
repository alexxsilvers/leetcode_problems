package main

import "log"

// Given the root of a binary tree, find the maximum average value of any subtree of that tree.
//
//(A subtree of a tree is any node of that tree plus all its descendants.
// The average value of a tree is the sum of its values, divided by the number of nodes.)
func main() {
	ex1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	log.Println(maximumAverageSubtree(ex1)) // 6

	ex2 := &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val: 1,
		},
	}
	log.Println(maximumAverageSubtree(ex2)) // 1.5

	ex3 := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	log.Println(maximumAverageSubtree(ex3)) // 2

	ex4 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 1},
	}
	log.Println(maximumAverageSubtree(ex4)) // 2
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Avg struct {
	Total float64
	Cnt   float64
}

func maximumAverageSubtree(root *TreeNode) float64 {
	var max float64 = 0
	dfs(root, &max)
	return max
}

func dfs(root *TreeNode, max *float64) *Avg {
	if root == nil {
		return &Avg{0, 0}
	}

	left := dfs(root.Left, max)
	right := dfs(root.Right, max)

	total := left.Total + right.Total + float64(root.Val)
	cnt := left.Cnt + right.Cnt + 1
	t := total / cnt
	if t > *max {
		*max = t
	}
	return &Avg{Total: total, Cnt: cnt}
}
