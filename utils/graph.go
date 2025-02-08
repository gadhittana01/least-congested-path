package utils

import (
	"container/heap"
)

type Edge struct {
	To         string
	Congestion int
}

type Graph struct {
	nodes map[string][]Edge
}

type IGraph interface {
	AddEdge(from, to string, congestion int)
	FindLeastCongestedPath(start, end string) string
	GetNode(nodeName string) []Edge
}

func NewGraph() IGraph {
	return &Graph{nodes: make(map[string][]Edge)}
}

func (g *Graph) GetNode(nodeName string) []Edge {
	return g.nodes[nodeName]
}

func (g *Graph) AddEdge(from, to string, congestion int) {
	g.nodes[from] = append(g.nodes[from], Edge{To: to, Congestion: congestion})
	g.nodes[to] = append(g.nodes[to], Edge{To: from, Congestion: congestion}) // Undirected
}

func (g *Graph) FindLeastCongestedPath(start, end string) string {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Node{start, 0, []string{start}})

	visited := make(map[string]bool)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)

		if current.name == end {
			return formatPath(current.path)
		}

		if visited[current.name] {
			continue
		}

		visited[current.name] = true
		for _, edge := range g.nodes[current.name] {
			if !visited[edge.To] {
				heap.Push(pq, &Node{
					name:            edge.To,
					totalCongestion: current.totalCongestion + edge.Congestion,
					path:            append(append([]string{}, current.path...), edge.To),
				})
			}
		}
	}
	return "No Path Found"
}

func formatPath(path []string) string {
	res := ""
	for i, p := range path {
		if i == 0 {
			res += p
		} else {
			res += " -> " + p
		}
	}
	return res
}
