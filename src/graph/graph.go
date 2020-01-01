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

// AddDirectedEdge adds an edge to the graph
func (g *Graph) AddDirectedEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.lock.Unlock()
}

// AddUndirectedEdge adds an edge to the graph
func (g *Graph) AddUndirectedEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
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

// MinimumDistance finds the shortest path
func (g *Graph) MinimumDistance(start, finish string) int {

	startNode := g.FindNode(Item(start))
	var nextLevelNodes []*Node = g.edges[*startNode]
	var visitedNodes = make(map[*Node]bool)

	// this is perhaps a weird algorithm. I basically separate the
	// graph into levels and return the level I find the target in!
	// There is bound to be a better way!
	for level := 1; len(nextLevelNodes) > 0; level++ {
		var nextNextLevelNodes []*Node = nil
		for _, node := range nextLevelNodes {
			if string(node.value) == finish {
				return level
			}
			visitedNodes[node] = true
			for _, child := range g.edges[*node] {
				if visitedNodes[child] {
					continue
				}
				nextNextLevelNodes = append(nextNextLevelNodes, child)
			}
		}
		nextLevelNodes = nextNextLevelNodes
	}

	return -1
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
