package main

import (
	"log"
	"strings"
)

// Given an absolute path for a file (Unix-style), simplify it. Or in other words,
// convert it to the canonical path.
// In a UNIX-style file system, a period . refers to the current directory.
// Furthermore, a double period .. moves the directory up a level. For more information,
// see: Absolute path vs relative path in Linux/Unix
// Note that the returned canonical path must always begin with a slash /,
// and there must be only a single slash / between two directory names.
// The last directory name (if it exists) must not end with a trailing /. Also,
// the canonical path must be the shortest string representing the absolute path.
func main() {
	log.Println(simplifyPath("/home/"))                // "/home"
	log.Println(simplifyPath("/../"))                  // "/"
	log.Println(simplifyPath("/home//foo/"))           // "/home/foo"
	log.Println(simplifyPath("/a/./b/../../c/"))       // "/c"
	log.Println(simplifyPath("/a/../../b/../c//.//"))  // "/c"
	log.Println(simplifyPath("/a//b////c/d//././/..")) // "/a/b/c"
}

func simplifyPath(path string) string {
	ignore := map[string]bool{
		".": true,
		"":  true,
	}

	st := newStack(len(path))
	paths := strings.Split(path, "/")
	for _, p := range paths {
		if p == ".." {
			st.Pop()
		} else if !ignore[p] {
			st.Push(p)
		}
	}

	if st.IsEmpty() {
		return "/"
	}

	ret := make([]string, 0)
	for !st.IsEmpty() {
		ret = append(ret, st.Pop())
	}

	sb := strings.Builder{}
	for i := len(ret) - 1; i >= 0; i-- {
		sb.WriteString("/")
		sb.WriteString(ret[i])
	}

	return sb.String()
}

type stack struct {
	buf  []string
	size int
	top  int
}

func newStack(size int) stack {
	return stack{
		buf:  make([]string, size),
		size: size,
		top:  -1,
	}
}
func (s *stack) IsEmpty() bool {
	return s.top == -1
}
func (s *stack) Push(str string) {
	if s.top+1 >= s.size {
		return
	}

	s.top++
	s.buf[s.top] = str
}
func (s *stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}

	v := s.buf[s.top]
	s.top--
	return v
}
func (s *stack) Top() string {
	if s.IsEmpty() {
		return ""
	}

	return s.buf[s.top]
}
