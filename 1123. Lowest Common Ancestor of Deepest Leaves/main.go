package main

import "log"

// Given a rooted binary tree, return the lowest common ancestor of its deepest leaves.
//
// Recall that:
//
// The node of a binary tree is a leaf if and only if it has no children
// The depth of the root of the tree is 0, and if the depth of a node is d,
// the depth of each of its children is d+1.
// The lowest common ancestor of a set S of nodes is the node A with the
// largest depth such that every node in S is in the subtree with root A.
func main() {
	ex1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	log.Println(lcaDeepestLeaves(ex1))

	ex2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3},
	}
	log.Println(lcaDeepestLeaves(ex2))

	ex3 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}
	log.Println(lcaDeepestLeaves(ex3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	node, _ := dfs(root)
	return node
}

func dfs(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	left, leftDepth := dfs(root.Left)
	right, rightDepth := dfs(root.Right)
	if leftDepth > rightDepth {
		return left, leftDepth + 1
	}
	if leftDepth < rightDepth {
		return right, rightDepth + 1
	}
	return root, leftDepth + 1
}
