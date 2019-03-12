package main

import "fmt"

func main() {
	circularQueue := Constructor(3)
	fmt.Println(circularQueue.EnQueue(1))  // return true
	fmt.Println(circularQueue.EnQueue(2))  // return true
	fmt.Println(circularQueue.EnQueue(3))  // return true
	fmt.Println(circularQueue.EnQueue(4))  // return false, the queue is full
	fmt.Println(circularQueue.Rear())  // return 3
	fmt.Println(circularQueue.IsFull())  // return true
	fmt.Println(circularQueue.DeQueue())  // return true
	fmt.Println(circularQueue.EnQueue(4))  // return true
	fmt.Println(circularQueue.Rear())  // return 4
}

type CircularQueue interface {
	/** Insert an element into the circular queue. Return true if the operation is successful. */
	EnQueue(value int) bool
	/** Delete an element from the circular queue. Return true if the operation is successful. */
	DeQueue() bool
	/** Get the front item from the queue. */
	Front() int
	/** Get the last item from the queue. */
	Rear() int
	/** Checks whether the circular queue is empty or not. */
	IsEmpty() bool
	/** Checks whether the circular queue is full or not. */
	IsFull() bool
}

type MyCircularQueue struct {
	data  []int
	front int
	rear  int
	size  int
}

func Constructor(size int) MyCircularQueue {
	return MyCircularQueue{
		data:  make([]int, size),
		front: -1,
		rear:  -1,
		size:  size,
	}
}

func (cq *MyCircularQueue) EnQueue(value int) bool {

}

func (cq *MyCircularQueue) DeQueue() bool {

}

func (cq *MyCircularQueue) Front() int {

}

func (cq *MyCircularQueue) Rear() int {

}

func (cq *MyCircularQueue) IsEmpty() bool {
}

func (cq *MyCircularQueue) IsFull() bool {

}
