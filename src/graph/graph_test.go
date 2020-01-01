package graph

import (
	"testing"
)

func fillGraph(g *Graph) {
	nA := g.AddOrFindNode("A")
	nB := g.AddOrFindNode("B")
	nC := g.AddOrFindNode("C")
	nD := g.AddOrFindNode("D")
	nE := g.AddOrFindNode("E")
	nF := g.AddOrFindNode("F")

	g.AddEdge(nA, nB)
	g.AddEdge(nA, nC)
	g.AddEdge(nB, nE)
	g.AddEdge(nC, nE)
	g.AddEdge(nE, nF)
	g.AddEdge(nD, nA)
}

func TestAdd(t *testing.T) {
	var g Graph
	fillGraph(&g)
	g.String()
}

func TestAddOrFind(t *testing.T) {
	var g Graph
	fillGraph(&g)

	var got = g.CountNodes()
	if got != 6 {
		t.Errorf("g.CountNodes() = %d not 6", got)
	}

	g.AddOrFindNode("A")
	got = g.CountNodes()
	if got != 6 {
		t.Errorf("g.CountNodes() = %d not 6", got)
	}

	g.AddNode("A")
	got = g.CountNodes()
	if got != 7 {
		t.Errorf("g.CountNodes() = %d not 7", got)
	}
}
