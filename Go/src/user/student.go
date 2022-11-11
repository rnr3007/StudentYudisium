package user

import (
	"fmt"
	"main/college/course"
	"main/util/input"
	"strings"
)

type Student struct {
	studentName         string
	studentCourseAmount int
	studentCourses      []course.Course
	studentSksTotal     int
	studentGpa          float32
	studentPredicate    string
}

type StudentDo interface {
	StudentName()
	StudentCourseAmount()
	StudentCourses()
	StudentSksTotal()
	StudentGpa()
	studentPredicate()
}

func (stdnt *Student) NewStudent() Student {
	stdnt.inputStudentName()
	stdnt.inputStudentCourseAmount()
	stdnt.inputStudentCourses()
	return Student{stdnt.studentName, stdnt.studentCourseAmount, stdnt.studentCourses, stdnt.studentSksTotal, stdnt.studentGpa, stdnt.studentPredicate}
}

func (stdnt *Student) inputStudentName() {
	fmt.Printf("\nYour name: ")
	stdnt.studentName = input.SafeAlphabetString()
}

func (stdnt *Student) inputStudentCourseAmount() {
	fmt.Printf("\nMany courses do you have: ")
	for isOk := true; isOk; isOk = (stdnt.studentCourseAmount < 2 || stdnt.studentCourseAmount > 12) {
		stdnt.studentCourseAmount = input.SafeInteger()
		if stdnt.studentCourseAmount < 2 || stdnt.studentCourseAmount > 12 {
			fmt.Printf("Enter proper course amount (2 - 12): ")
		}
	}
}

func (stdnt *Student) inputStudentCourses() {
	for counter := 0; counter < stdnt.studentCourseAmount; counter++ {
		var crs course.Course
		crs = crs.NewCourse()

		stdnt.studentSksTotal += crs.CourseSks()
		stdnt.studentGpa += (crs.CourseDoubleMark() * float32(crs.CourseSks()))
		fmt.Printf("\n\n         ;;PRESS [ENTER] TO CONTINUE;;")
		input.InputValue()
		stdnt.studentCourses = append(stdnt.studentCourses, crs)
	}
	stdnt.studentGpa = stdnt.studentGpa / float32(stdnt.studentSksTotal)
	stdnt.setStudentPredicate()
}

func (stdnt Student) StudentName() string {
	return stdnt.studentName
}

func (stdnt Student) StudentCourseAmount() int {
	return stdnt.studentCourseAmount
}

func (stdnt Student) StudentCourses() []course.Course {
	return stdnt.studentCourses
}

func (stdnt Student) StudentSksTotal() int {
	return stdnt.studentSksTotal
}

func (stdnt Student) StudentGpa() float32 {
	return stdnt.studentGpa
}

func (stdnt Student) StudentPredicate() string {
	return stdnt.studentPredicate
}

func (stdnt Student) setStudentPredicate() {
	for _, crs := range stdnt.studentCourses {
		if strings.EqualFold(crs.CourseLetterMark(), "E") {
			stdnt.studentPredicate = "PENDING"
			return
		}
	}

	if stdnt.studentGpa >= 3.5 {
		stdnt.studentPredicate = "CUM LAUDE"
	} else if stdnt.studentGpa >= 2.75 && stdnt.studentGpa < 3.5 {
		stdnt.studentPredicate = "VERY SATISFYING"
	} else if stdnt.studentGpa >= 2 && stdnt.studentGpa < 2.75 {
		stdnt.studentPredicate = "SATISFYING"
	}
}
