package model

type GroupingStrategy interface {
	Group(students []*Student) []*Group
}
