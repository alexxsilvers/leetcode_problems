package main

import "github.com/davecgh/go-spew/spew"

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 9}

	resultList := addTwoNumbers(l1, l2)
	spew.Dump(resultList) // 7 > 0 > 8
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultList := &ListNode{}
	head := resultList
	digit := 0
	for l1 != nil || l2 != nil {
		sum := digit
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		head.Next = &ListNode{Val: sum % 10}
		head = head.Next
		digit = sum / 10
	}

	if digit > 0 {
		head.Next = &ListNode{Val: digit}
	}

	return resultList.Next
}
