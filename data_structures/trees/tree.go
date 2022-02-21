package main

import "fmt"

/**
			9
	4				20
1		6		15			170
*/

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

// String prints a visual representation of the tree
func (t *Tree) String() {
	fmt.Println("------------------------------------------------")
	stringify(t.root, 0)
	fmt.Println("------------------------------------------------")
}

// https://flaviocopes.com/golang-data-structure-binary-search-tree/ for String and stringify. I didn't have any more time to think 💀
// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.value)
		stringify(n.right, level)
	}
}

func (t *Tree) Insert(value int) {
	n := &Node{}
	n.value = value
	if t.root == nil {
		t.root = n
	} else {
		currentNode := t.root
		for {
			if value < currentNode.value {
				// Go left
				if currentNode.left == nil {
					currentNode.left = n
					return
				}
				currentNode = currentNode.left
			} else {
				// Go right
				if currentNode.right == nil {
					currentNode.right = n
					return
				}
				currentNode = currentNode.right
			}
		}
	}
}

func (t *Tree) Lookup(value int) *Node {
	if t.root == nil {
		return nil
	}

	currentNode := t.root
	for currentNode != nil {
		if value < currentNode.value {
			// Check left
			currentNode = currentNode.left
		} else if value > currentNode.value {
			// Check right
			currentNode = currentNode.right
		} else if value == currentNode.value {
			return currentNode
		}
	}
	fmt.Println(currentNode)
	return currentNode
}

func (t *Tree) Remove(value int) *Node {
	if t.root == nil {
		return nil
	}

	currentNode := t.root
	var parentNode *Node
	for currentNode != nil {
		if value < currentNode.value {
			// Check left
			parentNode = currentNode
			currentNode = currentNode.left
		} else if value > currentNode.value {
			// Check right
			parentNode = currentNode
			currentNode = currentNode.right
		} else if value == currentNode.value {
			if currentNode.right == nil {
				// 1. No right child
				if parentNode == nil {
					t.root = currentNode.left
				} else {
					if currentNode.value < parentNode.value {
						// If parent > current value, current left child = child of parent
						parentNode.left = currentNode.left
					} else if currentNode.value > parentNode.value {
						// If parent < current value, current left child = child of parent
						parentNode.right = currentNode.left
					}
				}
			} else if currentNode.right.left == nil {
				// 2. Right child with no left child
				if parentNode == nil {
					t.root = currentNode.left
				} else {
					currentNode.right.left = currentNode.left
					if currentNode.value < parentNode.value {
						// If parent > current value, current right child = child of parent
						parentNode.left = currentNode.right
					} else if currentNode.value > parentNode.value {
						// If parent < current value, current right child = child of parent
						parentNode.right = currentNode.right
					}
				}
			} else {
				// 3. Right child with a left child
				// Check the right child's left most child
				leftMost := currentNode.right.left
				leftMostParent := currentNode.right
				for leftMost.left != nil {
					leftMostParent = leftMost
					leftMost = leftMost.left
				}

				// Paren't left subtree is now leftmost right subtree
				leftMostParent.left = leftMost.right
				leftMost.left = currentNode.left
				leftMost.right = currentNode.right

				if parentNode == nil {
					t.root = leftMost
				} else {
					if currentNode.value < parentNode.value {
						parentNode.left = leftMost
					} else if currentNode.value > parentNode.value {
						parentNode.right = leftMost
					}
				}
			}
			return currentNode
		}
	}
	fmt.Println(currentNode)
	return currentNode
}

func main() {
	t := &Tree{}
	t.Insert(9)
	t.Insert(4)
	t.Insert(6)
	t.Insert(20)
	t.Insert(170)
	t.Insert(15)
	t.Insert(1)
	t.Remove(1)
	t.Lookup(9)
	t.String()
}
