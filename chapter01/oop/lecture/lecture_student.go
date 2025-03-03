package lecture

type LectureAttendance struct {
	lecture *Lecture
	student *Student
	Grade   int
}

func (a *LectureAttendance) SetGrade(grade int) {
	if grade < 0 || grade > 100 {
		return
	}
	a.Grade = grade
}
