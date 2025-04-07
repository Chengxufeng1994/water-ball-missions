package adpater

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/relationship_analyzer/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/relationship_analyzer/internal/thirdparty"
)

type RelationshipAnalyzerAdapter struct {
	analyzer     thirdparty.SuperRelationshipAnalyzer
	relationship map[string][]string
}

var _ domain.RelationshipAnalyzerV1 = (*RelationshipAnalyzerAdapter)(nil)

func NewRelationshipAnalyzerAdapter(analyzer thirdparty.SuperRelationshipAnalyzer) *RelationshipAnalyzerAdapter {
	return &RelationshipAnalyzerAdapter{
		analyzer:     analyzer,
		relationship: make(map[string][]string),
	}
}

func (r *RelationshipAnalyzerAdapter) Parse(script string) {
	convScript := r.convertScript(script)
	r.analyzer.Init(convScript)
}

func (r *RelationshipAnalyzerAdapter) GetMutualFriends(name1, name2 string) []string {
	var commonFriends []string
	for _, friend := range r.relationship[name1] {
		if ok := r.analyzer.IsMutualFriends(friend, name1, name2); ok {
			commonFriends = append(commonFriends, friend)
		}
	}
	return commonFriends
}

func (r *RelationshipAnalyzerAdapter) convertScript(script string) string {
	rows := strings.SplitSeq(script, "\n")
	edges := make(map[string]struct{})
	for row := range rows {
		parts := strings.Split(row, ": ")
		node := parts[0]
		neighbors := strings.Split(parts[1], " ")
		for i := range neighbors {
			neighbor := neighbors[i]
			a, b := node, neighbor
			if a > b {
				a, b = b, a
			}
			key := fmt.Sprintf("%s -- %s", a, b)
			edges[key] = struct{}{}
		}

		r.relationship[node] = neighbors
	}
	// 收集並排序
	var sortedEdges []string
	for edge := range edges {
		sortedEdges = append(sortedEdges, edge)
	}
	sort.Strings(sortedEdges)

	return strings.Join(sortedEdges, "\n")
}
