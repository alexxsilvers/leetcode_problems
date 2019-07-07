package main

import "log"

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.
func main() {
	log.Println(isValid("()"))     // true
	log.Println(isValid("()[]{}")) // true
	log.Println(isValid("{]"))     // false
	log.Println(isValid("([)]"))   // false
	log.Println(isValid("{[]}"))   // true
	log.Println(isValid("{[()]}")) // true
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := NewStack(len(s))
	open := map[byte]struct{}{
		'{': struct{}{},
		'[': struct{}{},
		'(': struct{}{},
	}
	c := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	for i := 0; i < len(s); i++ {
		if _, exist := open[s[i]]; exist {
			stack.Push(s[i])
			continue
		}
		if cl, exist := c[s[i]]; exist {
			last := stack.Pop()
			if last != cl {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

type Stack struct {
	data []byte
	top  int
	size int
}

func NewStack(size int) *Stack {
	return &Stack{
		data: make([]byte, size),
		top:  -1,
		size: size,
	}
}
func (s *Stack) Push(v byte) bool {
	if s.top+1 >= s.size {
		return false
	}
	s.top++
	s.data[s.top] = v
	return true
}
func (s *Stack) Pop() byte {
	if s.IsEmpty() {
		return ' '
	}
	v := s.data[s.top]
	s.top--
	return v
}
func (s *Stack) Top() byte {
	if s.IsEmpty() {
		return ' '
	}
	return s.data[s.top]
}
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}
