package main

import "ds.com/ds/ds"

func main() {
	q := ds.NewQueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.PeekPrint()
}
