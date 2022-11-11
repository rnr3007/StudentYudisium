package course

import (
	"fmt"
	"main/util/input"
	"regexp"
)

type Course struct {
	courseName       string
	courseSks        int
	courseLetterMark string
	courseDoubleMark float32
}

type CourseDo interface {
	CourseName()
	CourseSks()
	CourseLetterMark()
	CourseDoubleMark()
}

func (crs *Course) NewCourse() Course {
	crs.inputCourseName()
	crs.inputCourseSks()
	crs.inputCourseLetterMark()
	return Course{crs.CourseName(), crs.CourseSks(), crs.CourseLetterMark(), crs.CourseDoubleMark()}
}

func (crs *Course) inputCourseName() {
	fmt.Printf("\nCourse name: ")
	crs.courseName = input.SafeAlphabetString()
}

func (crs *Course) inputCourseSks() {
	fmt.Printf("\nMany courseSks (2 - 6): ")
	for isOk := true; isOk; isOk = (crs.courseSks < 2 || crs.courseSks > 6) {
		crs.courseSks = input.SafeInteger()
		if crs.courseSks < 2 || crs.courseSks > 6 {
			fmt.Printf("Input the proper courseSks amount (2 - 6): ")
		}
	}
}

func (crs *Course) inputCourseLetterMark() {
	var regex, err = regexp.Compile("\\b[ABCDE]{1}\\b")
	if err == nil {
		for isOk := true; isOk; isOk = !regex.MatchString(crs.courseLetterMark) {
			fmt.Printf("\nCourse letter mark (A/B/C/D/E): ")
			crs.courseLetterMark = input.SafeAlphabetString()
			if !regex.MatchString(crs.courseLetterMark) {
				fmt.Printf("Input the proper letter mark (A/B/C/D/E)!\n")
			}
		}
	}
}

func (crs Course) CourseName() string {
	return crs.courseName
}

func (crs Course) CourseSks() int {
	return crs.courseSks
}

func (crs Course) CourseLetterMark() string {
	return crs.courseLetterMark
}

func (crs Course) CourseDoubleMark() float32 {
	switch crs.courseLetterMark {
	case "A":
		return 4.0
	case "B":
		return 3.0
	case "C":
		return 2.0
	case "D":
		return 1.0
	case "E":
		return 0.0
	}
	return 0
}
