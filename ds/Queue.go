package ds

import "fmt"

type Queue struct {
	first  *node
	last   *node
	length int
}

func NewQueue(val int) *Queue {
	n := node{value: val}

	return &Queue{first: &n, last: &n, length: 1}
}

func (q *Queue) Peek() *node {
	return q.first
}

func (q *Queue) Enqueue(val int) *Queue {
	n := node{value: val}
	if q.isEmpty() {
		q.first = &n
		q.last = &n
	} else {
		q.last.next = &n
		q.last = &n
	}
	q.length++

	return q
}

func (q *Queue) Dequeue() *Queue {
	if q.isEmpty() {
		return nil
	}
	if q.first == q.last {
		q.last = nil
	}
	q.first = q.first.next
	q.length--

	return q
}

func (s *Queue) isEmpty() bool {
	return s.length == 0
}

func (s *Queue) PeekPrint() {
	n := s.Peek()

	if n != nil {
		fmt.Println(n.value)
	} else {
		fmt.Println("nil")
	}
}
