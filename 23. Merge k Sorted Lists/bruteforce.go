package main

import "sort"

// walk throw each list and store values to slice
// then sort it and create new list
//
func mergeKListsBruteForce(lists []*ListNode) *ListNode {
	values := make([]int, 0)

	for _, list := range lists {
		head := list
		for head != nil {
			values = append(values, head.Val)
			head = head.Next
		}
	}

	sort.Ints(values)

	head := &ListNode{}
	curr := head
	for i := 0; i < len(values); i++ {
		curr.Next = &ListNode{Val: values[i]}
		curr = curr.Next
	}

	return head.Next
}
