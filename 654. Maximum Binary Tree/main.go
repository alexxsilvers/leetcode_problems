package main

import "github.com/davecgh/go-spew/spew"

// Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:
//
// The root is the maximum number in the array.
// The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
// The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
// Construct the maximum tree by the given array and output the root node of this tree.
func main() {
	spew.Dump(constructMaximumBinaryTree([]int{3,2,1,6,0,5}))
	//spew.Dump(constructMaximumBinaryTree([]int{4,5,10,6,-10}))
	//spew.Dump(constructMaximumBinaryTree([]int{5, 10}))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return construct(nums)
}

func construct(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxInd, maxVal := 0, nums[0]
	for i, n := range nums {
		if n > maxVal {
			maxInd, maxVal = i, n
		}
	}

	node := &TreeNode{
		Val: maxVal,
	}

	node.Left = construct(nums[:maxInd])
	node.Right = construct(nums[maxInd+1:])

	return node
}
