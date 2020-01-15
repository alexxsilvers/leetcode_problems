package main

import (
	"github.com/davecgh/go-spew/spew"
)

// Merge two sorted linked lists and return it as a new list. The new list should be
// made by splicing together the nodes of the first two lists.
func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l3 := mergeTwoLists(l1, l2)
	spew.Dump(l3) // 1 -> 1 -> 2 -> 3 -> 4 -> 4
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	prev := preHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}

		prev = prev.Next
	}

	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}

	return preHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
