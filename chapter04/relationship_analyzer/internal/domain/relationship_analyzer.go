package domain

type RelationshipAnalyzerV1 interface {
	Parse(script string)
	GetMutualFriends(name1 string, name2 string) []string
}

type RelationshipAnalyzerV2 interface {
	Parse(script string) RelationshipGraph
}
