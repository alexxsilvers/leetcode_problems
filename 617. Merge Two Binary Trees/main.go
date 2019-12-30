package main

import "github.com/davecgh/go-spew/spew"

// Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two
// trees are overlapped while the others are not.
//
// You need to merge them into a new binary tree. The merge rule is that if two nodes overlap,
// then sum node values up as the new value of the merged node. Otherwise, the NOT null node will
// be used as the node of new tree.
func main() {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	t2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	resp := mergeTreesStack(t1, t2)
	spew.Dump(resp)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive approach
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)

	return t1
}

func mergeTreesStack(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	st := newStack()
	st.Push([]*TreeNode{t1, t2})
	for !st.IsEmpty() {
		t := st.Pop()
		if t == nil || t[0] == nil || t[1] == nil {
			continue
		}

		t[0].Val += t[1].Val

		if t[0].Left == nil {
			t[0].Left = t[1].Left
		} else {
			st.Push([]*TreeNode{t[0].Left, t[1].Left})
		}

		if t[0].Right == nil {
			t[0].Right = t[1].Right
		} else {
			st.Push([]*TreeNode{t[0].Right, t[1].Right})
		}
	}

	return t1
}

type stack struct {
	data [][]*TreeNode
}

func newStack() *stack {
	return &stack{
		data: make([][]*TreeNode, 0),
	}
}
func (s *stack) Push(nodes []*TreeNode) {
	s.data = append(s.data, nodes)
}
func (s *stack) Pop() []*TreeNode {
	if s.IsEmpty() {
		return nil
	}

	ind := len(s.data) - 1
	v := s.data[ind]

	s.data = s.data[:ind]

	return v
}
func (s *stack) Top() []*TreeNode {
	if s.IsEmpty() {
		return nil
	}

	return s.data[len(s.data)-1]
}
func (s *stack) IsEmpty() bool {
	return len(s.data) == 0
}
