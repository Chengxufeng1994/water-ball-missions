package lecture

import "slices"

type Lecture struct {
	Name              string
	LectureAttendance []*LectureAttendance
}

func NewLecture(name string) *Lecture {
	return &Lecture{
		Name: name,
	}
}

func (l *Lecture) SignUp(student *Student) {
	if l.CheckStudentIsSigUp(student) {
		return
	}
	attendance := &LectureAttendance{
		lecture: l,
		student: student,
	}
	l.LectureAttendance = append(l.LectureAttendance, attendance)
	student.AddLectureAttendance(attendance)
}

func (l *Lecture) SignOff(student *Student) {
	if !l.CheckStudentIsSigUp(student) {
		return
	}

	var flag bool
	var i int
	for i < len(l.LectureAttendance) {
		if l.LectureAttendance[i].student == student {
			flag = true
			break
		}
		i++
	}
	if !flag {
		return
	}
	l.LectureAttendance = slices.Delete(l.LectureAttendance, i, i+1)
	student.RemoveLectureAttendance(l.LectureAttendance[i])
}

func (l *Lecture) CheckStudentIsSigUp(student *Student) bool {
	for _, v := range l.LectureAttendance {
		if v.student == student {
			return true
		}
	}
	return false
}
