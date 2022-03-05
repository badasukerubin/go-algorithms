package trees

import "fmt"

/**
			9
	4				20
1		6		15			170
*/

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

// String prints a visual representation of the tree
func (t *Tree) String() {
	fmt.Println("------------------------------------------------")
	stringify(t.Root, 0)
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
		stringify(n.Left, level)
		fmt.Printf(format+"%d\n", n.Value)
		stringify(n.Right, level)
	}
}

func (t *Tree) Insert(value int) {
	n := &Node{}
	n.Value = value
	if t.Root == nil {
		t.Root = n
	} else {
		currentNode := t.Root
		for {
			if value < currentNode.Value {
				// Go left
				if currentNode.Left == nil {
					currentNode.Left = n
					return
				}
				currentNode = currentNode.Left
			} else {
				// Go right
				if currentNode.Right == nil {
					currentNode.Right = n
					return
				}
				currentNode = currentNode.Right
			}
		}
	}
}

func (t *Tree) Lookup(value int) *Node {
	if t.Root == nil {
		return nil
	}

	currentNode := t.Root
	for currentNode != nil {
		if value < currentNode.Value {
			// Check left
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			// Check right
			currentNode = currentNode.Right
		} else if value == currentNode.Value {
			return currentNode
		}
	}
	fmt.Println(currentNode)
	return currentNode
}

func (t *Tree) Remove(value int) *Node {
	if t.Root == nil {
		return nil
	}

	currentNode := t.Root
	var parentNode *Node
	for currentNode != nil {
		if value < currentNode.Value {
			// Check left
			parentNode = currentNode
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			// Check right
			parentNode = currentNode
			currentNode = currentNode.Right
		} else if value == currentNode.Value {
			if currentNode.Right == nil {
				// 1. No right child
				if parentNode == nil {
					t.Root = currentNode.Left
				} else {
					if currentNode.Value < parentNode.Value {
						// If parent > current value, current left child = child of parent
						parentNode.Left = currentNode.Left
					} else if currentNode.Value > parentNode.Value {
						// If parent < current value, current left child = child of parent
						parentNode.Right = currentNode.Left
					}
				}
			} else if currentNode.Right.Left == nil {
				// 2. Right child with no left child
				if parentNode == nil {
					t.Root = currentNode.Left
				} else {
					currentNode.Right.Left = currentNode.Left
					if currentNode.Value < parentNode.Value {
						// If parent > current value, current right child = child of parent
						parentNode.Left = currentNode.Right
					} else if currentNode.Value > parentNode.Value {
						// If parent < current value, current right child = child of parent
						parentNode.Right = currentNode.Right
					}
				}
			} else {
				// 3. Right child with a left child
				// Check the right child's left most child
				leftMost := currentNode.Right.Left
				leftMostParent := currentNode.Right
				for leftMost.Left != nil {
					leftMostParent = leftMost
					leftMost = leftMost.Left
				}

				// Paren't left subtree is now leftmost right subtree
				leftMostParent.Left = leftMost.Right
				leftMost.Left = currentNode.Left
				leftMost.Right = currentNode.Right

				if parentNode == nil {
					t.Root = leftMost
				} else {
					if currentNode.Value < parentNode.Value {
						parentNode.Left = leftMost
					} else if currentNode.Value > parentNode.Value {
						parentNode.Right = leftMost
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
