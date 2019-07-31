package main

import "github.com/davecgh/go-spew/spew"

// Return the root node of a binary search tree that matches the given preorder traversal.
//
// (Recall that a binary search tree is a binary tree where for every node, any descendant of
// node.left has a value < node.val, and any descendant of node.right has a value > node.val.
// Also recall that a preorder traversal displays the value of the node first, then traverses
// node.left, then traverses node.right.)
func main() {
	spew.Dump(bstFromPreorder([]int{3, 1, 2}))
	spew.Dump(bstFromPreorder([]int{8, 5, 1, 7, 10, 12}))
	spew.Dump(bstFromPreorder([]int{4, 2}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	return helper(preorder, 0, len(preorder))
}

func helper(nums []int, start, end int) *TreeNode {
	if start >= end {
		return nil
	}

	node := &TreeNode{Val: nums[start]}
	index := end
	for i := start + 1; i < end; i++ {
		if nums[i] > nums[start] {
			index = i
			break
		}
	}

	node.Left = helper(nums, start+1, index)
	node.Right = helper(nums, index, end)
	return node
}
