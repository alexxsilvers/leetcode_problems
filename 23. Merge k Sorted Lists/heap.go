package main

func mergeKListsHeap(lists []*ListNode) *ListNode {
	h := newHeap()

	for _, list := range lists {
		if list != nil {
			h.push(list)
		}
	}

	ret := &ListNode{}
	head := ret
	for !h.isEmpty() {
		node := h.pop()
		ret.Next = node
		ret = ret.Next

		if node.Next != nil {
			h.push(node.Next)
		}
	}

	return head.Next
}

type heap struct {
	buf []*ListNode
	top int
}

func newHeap() heap {
	return heap{
		buf: make([]*ListNode, 0),
		top: -1,
	}
}

func (h heap) parent(i int) int {
	return (i - 1) / 2
}
func (h heap) left(i int) int {
	return 2*i + 1
}
func (h heap) right(i int) int {
	return 2*i + 2
}
func (h *heap) push(node *ListNode) {
	h.buf = append(h.buf, node)
	h.top++

	i := h.top
	parent := h.parent(i)
	for i > 0 && h.buf[i].Val < h.buf[parent].Val {
		h.buf[i], h.buf[parent] = h.buf[parent], h.buf[i]

		i = parent
		parent = h.parent(i)
	}
}
func (h *heap) pop() *ListNode {
	v := h.buf[0]
	h.buf[0] = h.buf[h.top]
	h.buf = h.buf[:h.top]
	h.top--

	h.heapify(0)

	return v
}
func (h *heap) heapify(i int) {
	for {
		left := h.left(i)
		right := h.right(i)
		smallest := i

		if left <= h.top && h.buf[left].Val < h.buf[smallest].Val {
			smallest = left
		}

		if right <= h.top && h.buf[right].Val < h.buf[smallest].Val {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.buf[smallest], h.buf[i] = h.buf[i], h.buf[smallest]
		i = smallest
	}
}
func (h *heap) isEmpty() bool {
	return h.top == -1
}
