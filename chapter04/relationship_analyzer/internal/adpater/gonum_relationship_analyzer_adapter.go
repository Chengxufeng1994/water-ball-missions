package adpater

import (
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/relationship_analyzer/internal/domain"
	"gonum.org/v1/gonum/graph/simple"
)

// 具名節點對應的 ID 映射
type nodeMap map[string]int64

type GonumRelationshipAnalyzerAdapter struct {
	nextID int64
}

var _ domain.RelationshipAnalyzerV2 = (*GonumRelationshipAnalyzerAdapter)(nil)

func NewGonumRelationshipAnalyzerAdapter() *GonumRelationshipAnalyzerAdapter {
	return &GonumRelationshipAnalyzerAdapter{}
}
func (adapter *GonumRelationshipAnalyzerAdapter) Parse(script string) domain.RelationshipGraph {
	g := simple.NewUndirectedGraph()
	nodes := make(map[string]int64)
	lines := strings.Split(script, "\n")

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		from := strings.TrimSpace(parts[0])
		neighbors := strings.Split(parts[1], " ")
		fromID := getOrCreateNodeID(from, nodes, g, &adapter.nextID)

		for _, to := range neighbors {
			to = strings.TrimSpace(to)
			if to == "" {
				continue
			}
			toID := getOrCreateNodeID(to, nodes, g, &adapter.nextID)
			g.SetEdge(g.NewEdge(g.Node(fromID), g.Node(toID)))
		}
	}

	return NewGonumRelationshipGraphAdapter(g, nodes)
}

func getOrCreateNodeID(name string, nodes nodeMap, g *simple.UndirectedGraph, nextID *int64) int64 {
	if id, ok := nodes[name]; ok {
		return id
	}
	id := *nextID
	*nextID++
	nodes[name] = id
	g.AddNode(simple.Node(id))
	return id
}
