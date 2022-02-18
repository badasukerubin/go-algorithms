package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

// O(n)
func (l *LinkedList) Get() {
	if l.len == 0 {
		fmt.Print("Linked list has no nodes")
	}

	pl := l.head
	for i := 0; i < l.len; i++ {
		fmt.Printf("Node: %v", pl)
		pl = pl.next
	}
}

// O(1)
func (l *LinkedList) Append(value int) {
	n := &Node{}
	n.value = value

	if l.len == 0 {
		l.head = n
		l.len++
		return
	}

	last := l.head
	for {
		if last.next == nil {
			break
		}
		last = last.next
	}

	last.next = n
	l.len++
}

// O(1)
func (l *LinkedList) Prepend(value int) {
	n := &Node{}
	n.value = value
	n.next = l.head

	l.head = n
	l.len++
}

func (l *LinkedList) traverseToIndex(index int) *Node {
	counter := 0
	currentNode := l.head

	for counter != index {
		currentNode = currentNode.next
		counter++
	}

	return currentNode
}

// O(n)
func (l *LinkedList) Insert(index int, value int) {
	if index >= l.len {
		l.Append(value)
		return
	}

	n := &Node{}
	n.value = value
	leader := l.traverseToIndex(index - 1)
	holdingPointer := leader.next
	leader.next = n
	n.next = holdingPointer

	l.len++
}

// O(1)
func (l *LinkedList) Remove(index int) {
	leader := l.traverseToIndex(index - 1)
	unwantedNode := leader.next
	leader.next = unwantedNode.next

	l.len--
}

// O(n)
func (l *LinkedList) Reverse() {
	if l.len == 1 {
		return
	}
	first := l.head
	second := first.next

	for second != nil {
		temp := second.next
		second.next = first
		first = second
		second = temp
	}
	l.head.next = nil
	l.head = first
}

func main() {
	l := &LinkedList{}
	l.Append(10)
	l.Append(20)
	l.Prepend(0)
	l.Append(40)
	l.Insert(3, 30)
	l.Insert(2, 100)
	l.Remove(2)
	l.Get()
	l.Reverse()
	l.Get()
}
