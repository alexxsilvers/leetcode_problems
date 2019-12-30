package main

import "github.com/davecgh/go-spew/spew"

// You are given two non-empty linked lists representing two non-negative integers. The digits
// are stored in reverse order and each of their nodes contain a single digit. Add the two numbers
// and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 9}
	resultList := addTwoNumbers2(l1, l2)
	spew.Dump(resultList) // 0 -> 1

	l3 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l4 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	resultList1 := addTwoNumbers2(l3, l4)
	spew.Dump(resultList1) // 7 -> 0 -> 8

	l5 := &ListNode{Val: 2}
	l6 := &ListNode{Val: 8, Next: &ListNode{Val: 1}}
	resultList2 := addTwoNumbers2(l5, l6)
	spew.Dump(resultList2) // 0 -> 2

	l7 := &ListNode{Val: 9, Next: &ListNode{Val: 1}}
	l8 := &ListNode{Val: 1}
	resultList3 := addTwoNumbers2(l7, l8)
	spew.Dump(resultList3) // 0 -> 2

	l9 := &ListNode{Val: 1}
	l10 := &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	resultList4 := addTwoNumbers2(l9, l10)
	spew.Dump(resultList4) // 0 -> 0 -> 1
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime O(max(l1, l2)), Space O(max(l1, l2))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	current := dummy
	add := 0

	for l1 != nil || l2 != nil || add > 0 {
		v := add

		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}

		add = v / 10

		current.Next = &ListNode{Val: v % 10}
		current = current.Next
	}

	return dummy.Next
}

// Runtime O(max(l1, l2), Space O(1)
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	tail, head := new(ListNode), l1
	add := 0

	for l1 != nil {
		if l1.Next == nil {
			if l2 != nil {
				l1.Next = l2.Next
				l2.Next = nil
			}
		}

		v := add + l1.Val

		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}

		add = v / 10
		l1.Val = v % 10
		tail = l1
		l1 = l1.Next
	}

	if add > 0 {
		tail.Next = &ListNode{Val: add}
	}

	return head
}
