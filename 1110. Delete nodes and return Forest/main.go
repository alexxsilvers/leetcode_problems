package main

import "github.com/davecgh/go-spew/spew"

// Given the root of a binary tree, each node in the tree has a distinct value.
//
// After deleting all nodes with a value in to_delete, we are left with a forest (a disjoint union of trees).
//
// Return the roots of the trees in the remaining forest.  You may return the result in any order.

// INTUITION
// Solve tree problem with recursion first (C)

// EXPLANATION
// If node is root(has no parent) and is not deleted, when we will add it to the result

// COMPLEXITY
// Runtime O(n)
// Space O(n)
func main() {
	ex1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	ret := delNodes(ex1, []int{3, 5})
	spew.Dump(ret)
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	forest := NewForest()
	delMap := make(map[int]bool, len(to_delete))
	for _, d := range to_delete {
		delMap[d] = true
	}
	helper(delMap, forest, root, true)
	return *forest
}

func NewForest() *Forest {
	r := make([]*TreeNode, 0)
	forest := Forest(r)
	return &forest
}

type Forest []*TreeNode

func (f *Forest) Push(n *TreeNode) {
	*f = append(*f, n)
}

func helper(delMap map[int]bool, forest *Forest, n *TreeNode, isParent bool) *TreeNode {
	if n == nil {
		return nil
	}
	deleted := delMap[n.Val]
	if isParent && !deleted {
		forest.Push(n)
	}
	n.Left = helper(delMap, forest, n.Left, deleted)
	n.Right = helper(delMap, forest, n.Right, deleted)
	if deleted {
		return nil
	}
	return n
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
