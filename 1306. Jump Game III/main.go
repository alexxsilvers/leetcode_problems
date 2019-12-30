package main

import "log"

// Given an array of non-negative integers arr, you are initially positioned at start index of the array.
// When you are at index i, you can jump to i + arr[i] or i - arr[i], check if you can reach to any index
// with value 0.
// Notice that you can not jump outside of the array at any time.
func main() {
	log.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5)) // true
	log.Println(canReach([]int{4, 2, 3, 0, 3, 1}, 0))    // true
	log.Println(canReach([]int{3, 0, 2, 1, 2}, 2))       // false
	log.Println(canReach([]int{0, 1}, 1))                // true
}

func canReach(arr []int, start int) bool {
	l := len(arr)
	visited := make(map[int]bool)

	stack := NewStack()
	stack.Push(start)

	for !stack.IsEmpty() {
		x := stack.Pop()
		if arr[x] == 0 {
			return true
		}

		if visited[x] {
			continue
		}
		visited[x] = true

		left := x - arr[x]
		if inRange(left, 0, l) {
			stack.Push(left)
		}

		right := x + arr[x]
		if inRange(right, 0, l) {
			stack.Push(right)
		}
	}

	return false
}

func inRange(x, a, b int) bool { return x >= a && x < b }

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{data: make([]int, 0)}
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x
}
