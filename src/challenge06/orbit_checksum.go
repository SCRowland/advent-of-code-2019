package challenge06

import (
	"graph"
	"strings"
)

const NODES_TO_REMOVE = 2

func creatOrbitGraph(orbitMap string, f func(*graph.Graph, *graph.Node, *graph.Node)) *graph.Graph {
	orbits := strings.Split(orbitMap, "\n")

	var g graph.Graph

	for _, orbit := range orbits {
		elements := strings.Split(orbit, ")")
		orbitee := elements[0]
		orbiter := elements[1]
		nOrbitee := g.AddOrFindNode(orbitee)
		nOrbiter := g.AddOrFindNode(orbiter)
		f(&g, nOrbiter, nOrbitee)
	}

	return &g
}

func addDirectedEdge(g *graph.Graph, n1 *graph.Node, n2 *graph.Node) {
	g.AddDirectedEdge(n1, n2)
}

func addUnirectedEdge(g *graph.Graph, n1 *graph.Node, n2 *graph.Node) {
	g.AddUndirectedEdge(n1, n2)
}

// OrbitCountChecksum calculates checksum for orbit map
func OrbitCountChecksum(orbitMap string) int {
	g := creatOrbitGraph(orbitMap, addDirectedEdge)

	return g.CountPathSteps()
}

// MinimumOrbitalTransferCount calculates the minimum path between nodes, excluding the nodes
func MinimumOrbitalTransferCount(orbitMap string) int {
	g := creatOrbitGraph(orbitMap, addUnirectedEdge)

	return g.MinimumDistance("YOU", "SAN") - NODES_TO_REMOVE
}
