package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Given a linked list, rotate the list to the right by k places, where k is non-negative.
func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	printList(list)
	fmt.Println(2 % 3)
	printList(rotateRight(list, 2))
}

func printList(head *ListNode) {
	nodes := make([]string, 0)
	for head != nil {
		nodes = append(nodes, strconv.Itoa(head.Val))
		head = head.Next
	}

	fmt.Println(strings.Join(nodes, " -> ") + " -> nil")
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	l := 1

	for curr.Next != nil {
		l++
		curr = curr.Next
	}

	curr.Next = head

	k = k % l
	m := l-k-1

	for m > 0 {
		head = head.Next
		m--
	}

	newHead := head.Next
	head.Next = nil

	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}
