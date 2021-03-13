package ds

import "fmt"

type ArrayStack struct {
	array []int
}

func NewArrayStack(val int) *ArrayStack {
	s := make([]int, 0)
	s = append(s, val)

	return &ArrayStack{s}
}

func (s *ArrayStack) Peek() *int {
	if !s.isEmpty() {
		return &s.array[len(s.array)-1]
	}

	return nil

}

func (s *ArrayStack) Push(val int) *ArrayStack {
	s.array = append(s.array, val)

	return s
}

func (s *ArrayStack) Pop() *ArrayStack {
	if s.isEmpty() {
		return nil
	}

	s.array = s.array[:len(s.array)-1]

	return s
}

func (s *ArrayStack) PeekPrint() {
	n := s.Peek()

	if n != nil {
		fmt.Println(*n)
	} else {
		fmt.Println("nil")
	}
}

func (s *ArrayStack) isEmpty() bool {
	return len(s.array) == 0
}
