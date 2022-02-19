package main

import "fmt"

// This implementation uses LinkedList
type Node struct {
	value int
	next  *Node
}

type Queue struct {
	first  int
	last   *Node
	length int
}

func (q *Queue) Peek() {
	fmt.Print(q.first)
}

// O(1)
func (q *Queue) Enqueue(value int) {
	n := &Node{}
	n.value = value

	if q.length == 0 {
		q.first = n.value
		q.last = n
	} else {
		q.last = n
		q.last.next = n
	}

	q.length++
}

// O(1)
func (q *Queue) Dequeue() {
	if q.length == 0 {
		return
	}
	if q.first == q.last.value {
		q.last = nil
	}
	q.first = q.last.next.value
	q.length--
}

func main() {
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Dequeue()
	q.Peek()
}
