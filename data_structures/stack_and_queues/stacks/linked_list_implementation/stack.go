package main

import "fmt"

// This implementation uses LinkedList
type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top    *Node
	bottom int
	length int
}

// O(1)
func (s *Stack) Peek() {
	fmt.Println(s.top.value)
}

// O(1)
func (s *Stack) Push(value int) {
	n := &Node{}
	n.value = value

	if s.length == 0 {
		s.top = n
		s.bottom = n.value
	} else {
		holdingPointer := s.top
		s.top = n
		s.top.next = holdingPointer
	}

	s.length++
}

// O(1)
func (s *Stack) Pop() {
	n := &Node{}
	if s.top == nil {
		return
	}
	if s.length == 0 {
		s.bottom = n.value
	}
	s.top = s.top.next
	s.length--
}

func main() {
	s := &Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Pop()
	s.Peek()
}
