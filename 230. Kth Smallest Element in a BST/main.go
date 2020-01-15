package main

import "log"

// Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
// Note:
// You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
func main() {
	tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	log.Println(kthSmallest(tree, 1))          // 1
	log.Println(kthSmallestRecursive(tree, 1)) // 1
	log.Println(kthSmallest(tree, 2))          // 2
	log.Println(kthSmallestRecursive(tree, 2)) // 2

	tree = &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6}}
	log.Println(kthSmallest(tree, 3))          // 3
	log.Println(kthSmallestRecursive(tree, 3)) // 3
	log.Println(kthSmallest(tree, 4))          // 4
	log.Println(kthSmallestRecursive(tree, 4)) // 4
}

func kthSmallestRecursive(root *TreeNode, k int) int {
	ret := make([]int, 0)
	dfs(root, &ret)

	return ret[k-1]
}

func dfs(node *TreeNode, ret *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, ret)
	*ret = append(*ret, node.Val)
	dfs(node.Right, ret)
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	st := newStack()

	for root != nil || !st.IsEmpty() {
		for root != nil {
			st.Push(root)
			root = root.Left
		}

		root = st.Pop()
		k -= 1
		if k == 0 {
			break
		}

		root = root.Right
	}

	return root.Val
}

type stack struct {
	buf []*TreeNode
}

func newStack() stack {
	return stack{buf: make([]*TreeNode, 0)}
}
func (s *stack) IsEmpty() bool {
	return len(s.buf) == 0
}
func (s *stack) Push(node *TreeNode) {
	s.buf = append(s.buf, node)
}
func (s *stack) Pop() *TreeNode {
	if s.IsEmpty() {
		return nil
	}

	ind := len(s.buf) - 1
	v := s.buf[ind]
	s.buf = s.buf[:ind]
	return v
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
