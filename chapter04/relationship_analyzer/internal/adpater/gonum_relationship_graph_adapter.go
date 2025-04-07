package adpater

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/relationship_analyzer/internal/domain"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
)

type GonumRelationshipGraphAdapter struct {
	graph *simple.UndirectedGraph
	nodes nodeMap
}

var _ domain.RelationshipGraph = (*GonumRelationshipGraphAdapter)(nil)

func NewGonumRelationshipGraphAdapter(graph *simple.UndirectedGraph, nodes nodeMap) *GonumRelationshipGraphAdapter {
	return &GonumRelationshipGraphAdapter{
		graph: graph,
		nodes: nodes,
	}
}

func (g *GonumRelationshipGraphAdapter) HasConnection(name1, name2 string) bool {
	id1, ok1 := g.nodes[name1]
	id2, ok2 := g.nodes[name2]
	if !ok1 || !ok2 {
		return false
	}

	// 使用 BFS 找是否可達
	visited := make(map[int64]bool)
	bfs := traverse.BreadthFirst{}
	startNode := g.graph.Node(id1)
	bfs.Walk(g.graph, startNode, func(n graph.Node, d int) bool {
		visited[n.ID()] = true
		return false
	})

	return visited[id2]
}
