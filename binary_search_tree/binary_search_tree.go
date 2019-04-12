package main

import "github.com/davecgh/go-spew/spew"

type BinarySearchTree interface {
	Add(item int)
	Contains(item int) bool
	Remove(item int) bool
	PreOrderTraversal(func(node *Node))
	PostOrderTraversal(func(node *Node))
	InOrderTraversal(func(node *Node))
	Clear()
}

func main() {
	bst := BST{}
	bst.Add(5)
	bst.Add(7)
	bst.Add(3)
	bst.Add(13)
	bst.Add(6)
	bst.Add(10)

	spew.Dump(bst)
}

func NewNode(item int) *Node {
	return &Node{
		value: item,
	}
}

type Node struct {
	value int
	left *Node
	right *Node
}

func (n *Node) CompareTo(value int) int {
	if value == n.value {
		return 0
	}

	if n.value > value {
		return -1
	}

	return 1
}

type BST struct {
	head *Node
}

func (bst *BST) Add(item int) {
	if bst.head == nil {
		bst.head = NewNode(item)
	} else {
		bst.add(bst.head, item)
	}
}

// recursive add
func (bst *BST) add(node *Node, item int) {
	if node.CompareTo(item) < 0 {
		if node.left == nil {
			node.left = NewNode(item)
		} else {
			bst.add(node.left, item)
		}
	} else {
		if node.right == nil {
			node.right = NewNode(item)
		} else {
			bst.add(node.right, item)
		}
	}
}
