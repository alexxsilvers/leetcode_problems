package main

import "log"

// Given a binary tree, return the inorder traversal of its nodes values.
// (left, root, right)
func main() {
	tree := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}}}
	log.Println(inorderTraversal(tree))          // [1, 5, 3, 2]
	log.Println(inorderTraversalRecursive(tree)) // [1, 5, 3, 2]

	tree = &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	log.Println(inorderTraversal(tree))          // [1, 3, 2]
	log.Println(inorderTraversalRecursive(tree)) // [1, 3, 2]
}

func inorderTraversalRecursive(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, res)
	*res = append(*res, node.Val)
	dfs(node.Right, res)
}

func inorderTraversal(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	st := newStack()

	for root != nil || !st.IsEmpty() {
		for root != nil {
			st.Push(root)
			root = root.Left
		}

		root = st.Pop()
		res = append(res, root.Val)
		root = root.Right
	}

	return res
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
func (s *stack) IsEmpty() bool {
	return len(s.buf) == 0
}
func (s *stack) Pop() *TreeNode {
	if s.IsEmpty() {
		return nil
	}

	v := s.buf[len(s.buf)-1]
	s.buf = s.buf[:len(s.buf)-1]
	return v
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
