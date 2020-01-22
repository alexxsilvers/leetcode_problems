package main

import "log"

// Find the kth largest element in an unsorted array. Note that it is the kth
// largest element in the sorted order, not the kth distinct element.
func main() {
	log.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          // 5
	log.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) // 4
}

func findKthLargest(nums []int, k int) int {
	h := newHeap()

	for _, num := range nums {
		h.push(num)

		if h.size() > k {
			h.pop()
		}
	}

	return h.peek()
}

type heap struct {
	buf []int
	top int
}

func newHeap() heap {
	return heap{
		buf: make([]int, 0),
		top: -1,
	}
}
func (h *heap) size() int {
	return h.top + 1
}
func (h *heap) pop() int {
	v := h.buf[0]
	h.buf[0] = h.buf[h.top]
	h.buf = h.buf[:h.top]
	h.top--

	h.heapify(0)

	return v
}
func (h *heap) push(val int) {
	h.top++
	h.buf = append(h.buf, val)

	i := h.top
	parent := h.parent(i)
	for i > 0 && h.buf[i] < h.buf[parent] {
		h.buf[i], h.buf[parent] = h.buf[parent], h.buf[i]

		i = parent
		parent = h.parent(i)
	}
}
func (h *heap) peek() int {
	return h.buf[0]
}
func (h heap) parent(i int) int {
	return (i - 1) / 2
}
func (h heap) left(i int) int {
	return i*2 + 1
}
func (h heap) right(i int) int {
	return i*2 + 2
}
func (h *heap) heapify(i int) {
	for {
		left := h.left(i)
		right := h.right(i)
		largest := i

		if left <= h.top && h.buf[left] < h.buf[largest] {
			largest = left
		}

		if right <= h.top && h.buf[right] < h.buf[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		h.buf[largest], h.buf[i] = h.buf[i], h.buf[largest]
		i = largest
	}
}
