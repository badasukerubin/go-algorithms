package main

import "fmt"

type Node struct {
	value int
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

func (g *Graph) String() {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

func (g *Graph) AddVertex(node int) {
	n := &Node{}
	n.value = node
	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(node1, node2 int) {
	// Undirected graph
	n1 := &Node{}
	n2 := &Node{}
	n1.value = node1
	n2.value = node2

	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}

	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func main() {
	g := Graph{}
	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddEdge(3, 1)
	g.AddEdge(3, 4)
	g.AddEdge(4, 2)
	g.AddEdge(4, 5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 0)
	g.AddEdge(0, 2)
	g.AddEdge(6, 5)
	g.String()
}
