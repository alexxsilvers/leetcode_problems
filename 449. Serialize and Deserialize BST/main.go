package main

import (
	"github.com/davecgh/go-spew/spew"
)

// Serialization is converting a data structure or object into a sequence of
// bits so that it can be stored in a file or memory buffer, or transmitted
// across a network connection link to be reconstructed later in the same or
// another computer environment.
//
// Design an algorithm to serialize and deserialize a binary search tree. There
// is no restriction on how your serialization/deserialization algorithm should work.
// You need to ensure that a binary search tree can be serialized to a string,
// and this string can be deserialized to the original tree structure.
//
// The encoded string should be as compact as possible.
func main() {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val:  8,
				Left: nil,
				Right: &TreeNode{
					Val: 12,
				},
			},
		},
	}
	obj := Constructor()

	serialized := obj.serialize(tree)
	newRoot := obj.deserialize(serialized)
	spew.Dump(newRoot)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	values := []rune{}

	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		values = append(values, rune(node.Val))
		preOrder(node.Left)
		preOrder(node.Right)
	}

	preOrder(root)
	return string(values)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	var insert func(node *TreeNode, val int)
	insert = func(node *TreeNode, val int) {
		if val <= node.Val {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
			} else {
				insert(node.Left, val)
			}
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
			} else {
				insert(node.Right, val)
			}
		}
	}

	root := &TreeNode{}

	for i, n := range data {
		if i == 0 {
			root.Val = int(n)
		} else {
			insert(root, int(n))
		}
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
