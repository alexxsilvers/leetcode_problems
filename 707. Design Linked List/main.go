package main

func main() {

}

// MyLinkedList MyLinkedList
type MyLinkedList struct {
	size int
	head *ListNode
	tail *ListNode
}

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

// Constructor return new MyLinkedList
func Constructor() MyLinkedList {
	return MyLinkedList{
		size: 0,
	}
}

// Get return the value of the index-th node in the linked list. If the index is invalid, return -1.
func (mll *MyLinkedList) Get(index int) int {
	if mll.size == 0 || index < 0 {
		return -1
	}

	dummy := mll.head
	for dummy != nil {
		if index == 0 {
			return dummy.Val
		}

		index--
		dummy = dummy.Next
	}

	return -1
}

// AddAtHead add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (mll *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{
		Val:  val,
		Next: mll.head,
	}

	mll.head.Prev = node
	mll.head = node
	mll.size++
}

// AddAtTail append a node of value val to the last element of the linked list.
func (mll *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{
		Val:  val,
		Prev: mll.tail,
	}

	mll.tail.Next = node
	mll.tail = node
	mll.size++
}

// AddAtIndex Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will
// be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
func (mll *MyLinkedList) AddAtIndex(index int, val int) {
	if index > mll.size {
		return
	}

	if index == mll.size {
		node := &ListNode{
			Val:  val,
			Prev: mll.tail,
		}

		mll.tail.Next = node
		mll.tail = node
		mll.size++
		return
	}

	dummy := mll.head
	for dummy != nil {
		if index == 0 {
			node := &ListNode{
				Val:  val,
				Prev: dummy,
				Next: dummy.Next,
			}

			dummy.Next.Prev = node
			dummy.Next = node

			return
		}

		dummy = dummy.Next
		index--
	}
}

// DeleteAtIndex Delete the index-th node in the linked list, if the index is valid.
func (mll *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= mll.size {
		return
	}

	dummy := mll.head
	for dummy != nil {
		if index == 0 {
			dummy.Next.Prev = dummy.Prev
			dummy.Prev.Next = dummy.Next
			return
		}

		dummy = dummy.Next
		index--
	}

	return
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
