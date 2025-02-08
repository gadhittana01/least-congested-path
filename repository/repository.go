package repository

import "github.com/gadhittana01/least-congested-path/utils"

type Edge struct {
	to         string
	congestion int
}

type Graph struct {
	nodes map[string][]Edge
}

type IRepository interface {
	FindLeastCongestedPath(start, end string) string
	AddRoad(start, end string, congestion int)
}

type repository struct {
	graph utils.IGraph
}

func NewRepository(graph utils.IGraph) IRepository {
	return &repository{
		graph: graph,
	}
}

func (h *repository) FindLeastCongestedPath(start, end string) string {
	return h.graph.FindLeastCongestedPath(start, end)
}

func (h *repository) AddRoad(start, end string, congestion int) {
	h.graph.AddEdge(start, end, congestion)
}
