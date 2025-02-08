package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue_Len(t *testing.T) {
	pq := PriorityQueue{}
	assert.Equal(t, 0, pq.Len(), "Expected length to be 0 initially")

	node := &Node{name: "A", totalCongestion: 3}
	pq.Push(node)
	assert.Equal(t, 1, pq.Len(), "Expected length to be 1 after push")
}

func TestPriorityQueue_Swap(t *testing.T) {
	pq := PriorityQueue{
		&Node{name: "A", totalCongestion: 3},
		&Node{name: "B", totalCongestion: 1},
	}
	pq.Swap(0, 1)
	assert.Equal(t, "B", pq[0].name, "Expected B to be at index 0 after swap")
	assert.Equal(t, "A", pq[1].name, "Expected A to be at index 1 after swap")
}

func TestPriorityQueue_Push(t *testing.T) {
	pq := &PriorityQueue{}
	node := &Node{name: "A", totalCongestion: 3}
	pq.Push(node)
	assert.Equal(t, 1, pq.Len(), "Expected length to be 1 after push")
	assert.Equal(t, "A", (*pq)[0].name, "Expected node A to be at index 0 after push")
}

func TestPriorityQueue_Pop(t *testing.T) {
	pq := &PriorityQueue{
		&Node{name: "A", totalCongestion: 3},
		&Node{name: "B", totalCongestion: 1},
	}
	node := pq.Pop().(*Node)
	assert.Equal(t, "B", node.name, "Expected node B to be popped first")
	assert.Equal(t, 1, pq.Len(), "Expected length to be 1 after pop")
}
