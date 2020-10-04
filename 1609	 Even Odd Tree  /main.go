package main

import "log"

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   12,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:  9,
				Left: nil,
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	log.Println(isEvenOddTree(tree)) // true
}

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var bfs func(nodes []*TreeNode, level int, findedFalse *bool)
	bfs = func(nodes []*TreeNode, level int, findedFalse *bool) {
		if len(nodes) == 0 {
			return
		}

		if isEven(level) {
			if isEven(nodes[0].Val) {
				*findedFalse = true
				return
			}
			for i := 0; i+1 < len(nodes); i++ {
				if isEven(nodes[i].Val) || isEven(nodes[i+1].Val) {
					*findedFalse = true
					return
				}
				if nodes[i].Val >= nodes[i+1].Val {
					*findedFalse = true
					return
				}
			}
		} else { // odd
			if !isEven(nodes[0].Val) {
				*findedFalse = true
				return
			}
			for i := 0; i+1 < len(nodes); i++ {
				if !isEven(nodes[i].Val) || !isEven(nodes[i+1].Val) {
					*findedFalse = true
					return
				}
				if nodes[i].Val <= nodes[i+1].Val {
					*findedFalse = true
					return
				}
			}
		}

		newNodes := make([]*TreeNode, 0)
		for _, n := range nodes {
			if n.Left != nil {
				newNodes = append(newNodes, n.Left)
			}
			if n.Right != nil {
				newNodes = append(newNodes, n.Right)
			}
		}
		bfs(newNodes, level+1, findedFalse)
	}

	findedFalse := false
	bfs([]*TreeNode{root}, 0, &findedFalse)
	return !findedFalse
}

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
