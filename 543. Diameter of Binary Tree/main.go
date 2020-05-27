package main

func main() {
	// t := &TreeNode{}
}

// TreeNode tree node definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var diameter int
	maxDepth(&diameter, root)
	return diameter
}

func maxDepth(diameter *int, node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := maxDepth(diameter, node.Left)
	right := maxDepth(diameter, node.Right)
	*diameter = max(*diameter, left+right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
