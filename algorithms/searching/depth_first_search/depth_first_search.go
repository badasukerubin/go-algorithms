package main

import (
	"fmt"

	"github.com/badasukerubin/go-algorithms/data_structures/trees"
)

func DFSInOrder(node *trees.Node, list *[]int) []int {
	return traverseInOrder(node, list)
}

func DFSPreOrder(node *trees.Node, list *[]int) []int {
	return traversePreOrder(node, list)
}

func DFSPostOrder(node *trees.Node, list *[]int) []int {
	return traversePostOrder(node, list)
}

func traverseInOrder(node *trees.Node, list *[]int) []int {
	if node.Left != nil {
		traverseInOrder(node.Left, list)
	}

	*list = append(*list, node.Value)

	if node.Right != nil {
		traverseInOrder(node.Right, list)
	}

	return *list
}

func traversePreOrder(node *trees.Node, list *[]int) []int {
	*list = append(*list, node.Value)

	if node.Left != nil {
		traversePreOrder(node.Left, list)
	}

	if node.Right != nil {
		traversePreOrder(node.Right, list)
	}

	return *list
}

func traversePostOrder(node *trees.Node, list *[]int) []int {
	if node.Left != nil {
		traversePreOrder(node.Left, list)
	}

	if node.Right != nil {
		traversePreOrder(node.Right, list)
	}

	*list = append(*list, node.Value)

	return *list
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

	node := t.Root
	list := &[]int{}

	DFSInOrder := DFSInOrder(node, list)

	list = &[]int{}
	DFSPreOrder := DFSPreOrder(node, list)

	list = &[]int{}
	DFSPostOrder := DFSPostOrder(node, list)

	fmt.Println(DFSInOrder)
	fmt.Println(DFSPreOrder)
	fmt.Println(DFSPostOrder)
}
