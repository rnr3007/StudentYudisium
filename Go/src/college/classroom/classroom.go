package classroom

import (
	"main/user"
)

type Classroom struct {
	students []user.Student
}

type ClassroomDo interface {
	AddClassroomStudents()
	ClassroomStudents()
}

func (clsrm *Classroom) AddClassroomStudents() {
	var stdnt user.Student
	stdnt = stdnt.NewStudent()
	clsrm.students = append(clsrm.students, stdnt)
	clsrm.sortStudentByGpa()
}

func (clsrm Classroom) ClassroomStudents() []user.Student {
	return clsrm.students
}

func (clsrm Classroom) sortStudentByGpa() {
	var tempStdnt user.Student
	for index := len(clsrm.students) - 1; index > 0; index-- {
		for index2 := 0; index2 < index; index2++ {
			currentStdnt := clsrm.students[index]
			currentGpa := currentStdnt.StudentGpa()
			startStdnt := clsrm.students[index2]
			startGpa := startStdnt.StudentGpa()
			if currentGpa > startGpa {
				tempStdnt = currentStdnt
				clsrm.students[index] = startStdnt
				clsrm.students[index2] = tempStdnt
			}
		}
	}
}
