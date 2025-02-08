package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph_AddEdge(t *testing.T) {
	g := NewGraph()

	g.AddEdge("A", "B", 3)

	assert.Equal(t, 1, len(g.GetNode("A")), "Expected 1 edge for node A")
	assert.Equal(t, "B", g.GetNode("A")[0].To, "Expected edge from A to B")
	assert.Equal(t, 3, g.GetNode("A")[0].Congestion, "Expected congestion to be 3")

	assert.Equal(t, 1, len(g.GetNode("B")), "Expected 1 edge for node B")
	assert.Equal(t, "A", g.GetNode("B")[0].To, "Expected edge from B to A")
	assert.Equal(t, 3, g.GetNode("B")[0].Congestion, "Expected congestion to be 3")
}

func TestGraph_FindLeastCongestedPath(t *testing.T) {
	g := NewGraph()

	g.AddEdge("A", "B", 3)
	g.AddEdge("B", "C", 2)
	g.AddEdge("A", "C", 5)

	path := g.FindLeastCongestedPath("A", "C")
	assert.Equal(t, "A -> C", path, "Expected the least congested path to be A -> B -> C")

	path = g.FindLeastCongestedPath("A", "D")
	assert.Equal(t, "No Path Found", path, "Expected 'No Path Found' since D does not exist")
}
