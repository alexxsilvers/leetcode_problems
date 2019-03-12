package binary_search_tree

import (
	"fmt"
)

type node struct {
	data int
	left *node
	right *node
}

func NewNode(data int) *node {
	return &node{
		data: data,
	}
}

func Insert(root *node, data int) *node {
	if root == nil {
		return NewNode(data)
	}

	if data < root.data {
		root.left = Insert(root.left, data)
	} else if data > root.data {
		root.right = Insert(root.right, data)
	}

	return root
}

func Search(root *node, item int) int {
	if root == nil {
		return -1
	}

	if item == root.data {
		return root.data
	}

	if item < root.data {
		return Search(root.left, item)
	}

	if item > root.data {
		return Search(root.right, item)
	}

	return -1
}

func InOrder(root *node) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Printf("->%d", root.data)
	InOrder(root.right)
}
