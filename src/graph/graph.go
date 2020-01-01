package graph

import (
	"fmt"
	"sync"
)

// Item is an element in the graph
type Item string

// Node is a node in the graph
type Node struct {
	value Item
}

// Graph is an implementation of a graph datastructure
type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

// FindNode finds a node
func (g *Graph) FindNode(i Item) *Node {
	for _, node := range g.nodes {
		if node.value == i {
			return node
		}
	}
	return nil
}

// AddOrFindNode adds a node to the graph
func (g *Graph) AddOrFindNode(v string) *Node {
	n := g.FindNode(Item(v))
	if n != nil {
		return n
	}
	return g.AddNode(v)
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(v string) *Node {
	n := &Node{Item(v)}

	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()

	return n
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.lock.Unlock()
}

// CountNodes counts the nodes!
func (g *Graph) CountNodes() int {
	return len(g.nodes)
}

func (g *Graph) pathSteps(n *Node) int {
	var stepCount = 0

	parents := g.edges[*n]
	for _, parent := range parents {
		stepCount++
		stepCount += g.pathSteps(parent)
	}

	return stepCount
}

// CountPathSteps counts the total number of steps in each path
func (g *Graph) CountPathSteps() int {
	var stepCount = 0

	for i := 0; i < len(g.nodes); i++ {
		stepCount += g.pathSteps(g.nodes[i])
	}

	return stepCount
}

// String Renders the Graph
func (g *Graph) String() string {
	g.lock.RLock()
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
	g.lock.RUnlock()
	return s
}
