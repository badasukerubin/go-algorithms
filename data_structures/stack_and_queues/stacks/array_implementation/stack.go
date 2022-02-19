package main

import "fmt"

// This implementation uses arrays
type Stack struct {
	array []int
}

// O(1)
func (s *Stack) Peek() {
	fmt.Println(s.array[len(s.array)-1])
}

// O(1)
func (s *Stack) Push(value int) {
	s.array = append(s.array, value)
}

// O(1)
func (s *Stack) Pop() {
	s.array = (s.array)[:len(s.array)-1]
}

func main() {
	s := &Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Pop()
	s.Peek()
}
