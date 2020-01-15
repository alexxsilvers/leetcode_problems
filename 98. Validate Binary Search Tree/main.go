package main

import "log"

// Given a binary tree, determine if it is a valid binary search tree (BST).
// Assume a BST is defined as follows:
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
func main() {
	tree := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	log.Println(isValidBST(tree))          // true
	log.Println(isValidBSTIterative(tree)) // true

	tree = &TreeNode{Val: 2, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3}}
	log.Println(isValidBST(tree))          // false
	log.Println(isValidBSTIterative(tree)) // false

	tree = &TreeNode{Val: 10, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 20}}}
	log.Println(isValidBST(tree))          // false
	log.Println(isValidBSTIterative(tree)) // false

	log.Println(isValidBST(nil))          // true
	log.Println(isValidBSTIterative(nil)) // true

	tree = &TreeNode{}
}

func isValidBSTIterative(root *TreeNode) bool {
	var prev *TreeNode
	st := newStack()
	for root != nil || !st.IsEmpty() {
		for root != nil {
			st.Push(root)
			root = root.Left
		}

		root = st.Pop()
		if prev != nil && root.Val <= prev.Val {
			return false
		}
		prev = root
		root = root.Right
	}

	return true
}

func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(node *TreeNode, lower *int, upper *int) bool {
	if node == nil {
		return true
	}

	if lower != nil && node.Val <= *lower {
		return false
	}
	if upper != nil && node.Val >= *upper {
		return false
	}

	if !validate(node.Right, &node.Val, upper) {
		return false
	}
	if !validate(node.Left, lower, &node.Val) {
		return false
	}

	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	buf []*TreeNode
}

func newStack() stack {
	return stack{buf: make([]*TreeNode, 0)}
}
func (s *stack) Push(node *TreeNode) {
	s.buf = append(s.buf, node)
}
func (s *stack) Pop() *TreeNode {
	v := s.buf[len(s.buf)-1]
	s.buf = s.buf[:len(s.buf)-1]
	return v
}
func (s *stack) IsEmpty() bool {
	return len(s.buf) == 0
}
