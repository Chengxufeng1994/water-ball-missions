package thirdparty

import (
	"slices"
	"strings"
)

type SuperRelationshipAnalyzer struct {
	relationship map[string][]string
}

func NewSuperRelationshipAnalyzer() *SuperRelationshipAnalyzer {
	return &SuperRelationshipAnalyzer{
		relationship: make(map[string][]string),
	}
}

func (s *SuperRelationshipAnalyzer) Init(script string) {
	rows := strings.SplitSeq(script, "\n")
	for row := range rows {
		names := strings.Split(row, " -- ")
		if len(names) < 2 {
			continue
		}

		name1 := names[0]
		name2 := names[1]
		if s.relationship[name1] == nil {
			s.relationship[name1] = []string{}
		}
		s.relationship[name1] = append(s.relationship[name1], name2)

		if s.relationship[name2] == nil {
			s.relationship[name2] = []string{}
		}
		s.relationship[name2] = append(s.relationship[name2], name1)
	}
}

func (s *SuperRelationshipAnalyzer) IsMutualFriends(targetName, name1, name2 string) bool {
	friendsOfFriend1, exists := s.relationship[name1]
	if !exists {
		return false
	}

	friendsOfFriend2, exists := s.relationship[name2]
	if !exists {
		return false
	}

	return slices.Contains(friendsOfFriend1, targetName) && slices.Contains(friendsOfFriend2, targetName)
}
