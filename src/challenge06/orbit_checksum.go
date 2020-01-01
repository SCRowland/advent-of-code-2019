package challenge06

import (
	"graph"
	"strings"
)

func creatOrbitGraph(orbitMap string) *graph.Graph {
	orbits := strings.Split(orbitMap, "\n")

	var g graph.Graph

	for _, orbit := range orbits {
		elements := strings.Split(orbit, ")")
		orbitee := elements[0]
		orbiter := elements[1]
		nOrbitee := g.AddOrFindNode(orbitee)
		nOrbiter := g.AddOrFindNode(orbiter)
		g.AddEdge(nOrbiter, nOrbitee)
	}

	return &g
}

// OrbitCountChecksum calculates checksum for orbit map
func OrbitCountChecksum(orbitMap string) int {
	g := creatOrbitGraph(orbitMap)

	return g.CountPathSteps()
}
