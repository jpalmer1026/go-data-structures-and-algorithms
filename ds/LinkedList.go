package ds

import "fmt"

type LinkedList struct {
	head   *node
	tail   *node
	length int
}

func NewLinkedList(value int) *LinkedList {
	head := node{value: value, next: nil, previous: nil}
	l := LinkedList{head: &head, tail: &head, length: 1}

	return &l
}

func (ll *LinkedList) Append(value int) *LinkedList {
	newNode := node{value: value, next: nil, previous: ll.tail}
	ll.tail.next = &newNode
	ll.tail = &newNode
	ll.length++

	return ll
}

func (ll *LinkedList) Prepend(value int) *LinkedList {
	newNode := node{value: value, next: ll.head, previous: nil}
	ll.head.previous = &newNode
	ll.head = &newNode
	ll.length++

	return ll
}

// 1 -> 2 -> 3
func (ll *LinkedList) Insert(index, value int) *LinkedList {
	if index >= ll.length {
		return ll.Append(value)
	}

	if index == 0 {
		return ll.Prepend(value)
	}

	if index == ll.length-1 {
		return ll.Append(value)
	}

	newNode := node{value: value, next: nil, previous: nil}

	leader := ll.traverseToIndex(index - 1)
	follower := leader.next
	leader.next = &newNode
	newNode.previous = leader
	newNode.next = follower
	follower.previous = &newNode
	ll.length++

	return ll
}

func (ll *LinkedList) Remove(index int) *LinkedList {
	if index == 0 {
		ll.head = ll.head.next
		ll.head.previous = nil
		ll.length--

		return ll
	}

	if index >= ll.length-1 {
		tailLeader := ll.traverseToIndex(ll.length - 3)
		newTail := tailLeader.next
		ll.tail = newTail
		ll.tail.next = nil
		ll.tail.previous = tailLeader
		ll.length--

		return ll
	}

	leader := ll.traverseToIndex(index - 1)
	next := leader.next
	finalNext := next.next
	leader.next = finalNext
	finalNext.previous = leader
	ll.length--

	return ll
}

func (ll *LinkedList) Print() {
	for n := ll.head; n.next != nil; n = n.next {
		fmt.Print(n.value)
		fmt.Print("->")
	}
	fmt.Println(ll.tail.value)
}

func (ll *LinkedList) PrintReverse() {
	for n := ll.tail; n.previous != nil; n = n.previous {
		fmt.Print(n.value)
		fmt.Print("->")
	}
	fmt.Println(ll.head.value)
}

func (ll *LinkedList) traverseToIndex(index int) *node {
	counter := 0
	currentNode := ll.head
	for counter != index {
		currentNode = currentNode.next
		counter++
	}

	return currentNode
}

// 1 -> 10 -> 16 -> 88
func (ll *LinkedList) Reverse() *LinkedList {
	if ll.length == 1 {
		return ll
	}
	ll.tail = ll.head

	first := ll.head
	second := first.next
	for second != nil {
		temp := second.next
		second.next = first
		first = second
		second = temp
	}
	ll.head.next = nil
	ll.head = first

	return ll
}
