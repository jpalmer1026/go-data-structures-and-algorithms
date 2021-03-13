package ds

import "fmt"

type Stack struct {
	top    *node
	bottom *node
	length int
}

func NewStack(val int) *Stack {
	n := node{value: val}

	return &Stack{top: &n, bottom: &n, length: 1}
}

func (s *Stack) Peek() *node {
	return s.top
}

func (s *Stack) Push(val int) *Stack {
	n := node{value: val}
	if s.isEmpty() {
		s.top = &n
		s.bottom = &n
	} else {
		oldTop := s.top
		s.top = &n
		s.top.next = oldTop
	}
	s.length++

	return s
}

func (s *Stack) Pop() *Stack {
	if s.isEmpty() {
		return nil
	}
	if s.top == s.bottom {
		s.bottom = nil
	}
	s.top = s.top.next
	s.length--

	return s
}

func (s *Stack) PeekPrint() {
	n := s.Peek()

	if n != nil {
		fmt.Println(n.value)
	} else {
		fmt.Println("nil")
	}
}

func (s *Stack) isEmpty() bool {
	return s.length == 0
}
