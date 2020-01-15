package main

import "github.com/davecgh/go-spew/spew"

// Reverse a singly linked list.
func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	l2 := reverseListRecursive(l1)

	spew.Dump(l2) // 4 -> 3 -> 2 -> 1 -> nil
}

// Runtime - O(n), Space - O(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		nextTmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTmp
	}

	return prev
}

// Runtime - O(n), Space - O(n)
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

type ListNode struct {
	Val  int
	Next *ListNode
}
