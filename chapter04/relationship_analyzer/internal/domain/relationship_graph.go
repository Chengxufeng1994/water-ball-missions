package domain

type RelationshipGraph interface {
	HasConnection(name1, name2 string) bool
}
