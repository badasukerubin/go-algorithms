package main

import (
	"fmt"

	"github.com/badasukerubin/go-algorithms/data_structures/trees"
)

func breadthFirstSearch(t *trees.Tree) []int {
	curentNode := t.Root
	list := []int{}
	queue := []*trees.Node{}

	queue = append(queue, curentNode)

	for len(queue) > 0 {
		curentNode, queue = queue[0], queue[1:]
		list = append(list, curentNode.Value)

		// Memory consuming - we keep adding to the queue
		if curentNode.Left != nil {
			queue = append(queue, curentNode.Left)
		}

		if curentNode.Right != nil {
			queue = append(queue, curentNode.Right)
		}
	}

	return list
}

func breadthFirstSearchRecursive(queue []*trees.Node, list []int) []int {
	var curentNode *trees.Node

	if len(queue) == 0 {
		return list
	}

	curentNode, queue = queue[0], queue[1:]
	list = append(list, curentNode.Value)
	if curentNode.Left != nil {
		queue = append(queue, curentNode.Left)
	}

	if curentNode.Right != nil {
		queue = append(queue, curentNode.Right)
	}
	return breadthFirstSearchRecursive(queue, list)
}

func main() {
	t := &trees.Tree{}
	t.Insert(9)
	t.Insert(4)
	t.Insert(6)
	t.Insert(20)
	t.Insert(170)
	t.Insert(15)
	t.Insert(1)

	breadthFirstSearch := breadthFirstSearch(t)

	queue := []*trees.Node{}
	queue = append(queue, t.Root)
	list := []int{}
	breadthFirstSearchRecursive := breadthFirstSearchRecursive(queue, list)

	fmt.Println(breadthFirstSearch)
	fmt.Println(breadthFirstSearchRecursive)
}
