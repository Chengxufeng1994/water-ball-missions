package model

var count int

type Group struct {
	Number   int
	Students []*Student
}

func NewGroup() *Group {
	count += 1
	return &Group{
		Number:   count,
		Students: make([]*Student, 0),
	}
}

func (group *Group) AddStudents(students ...*Student) {
	group.Students = append(group.Students, students...)
}

func (group *Group) Size() int {
	return len(group.Students)
}

func (group *Group) SplitGroupBySize(size int) []*Group {
	groups := make([]*Group, 0)
	studentSize := group.Size()
	for i := 0; i < studentSize; i += size {
		end := min(i+size, studentSize)
		subGroup := NewGroup()
		subGroup.AddStudents(group.Students[i:end]...)
		groups = append(groups, subGroup)
	}

	return groups
}

func (group *Group) MergeGroup(other *Group) {
	group.Students = append(group.Students, other.Students...)
}
