package main

import "github.com/davecgh/go-spew/spew"

// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
func main() {
	lists := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6}},
	}

	res := mergeKListsBruteForce(lists)

	spew.Dump(res)
	// 1 -> 1 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6
}

type ListNode struct {
	Val  int
	Next *ListNode
}
