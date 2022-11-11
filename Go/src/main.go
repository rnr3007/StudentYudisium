package main

import (
	"fmt"
	"main/appearance"
	"main/college/classroom"
	"main/util/input"
)

var clsrm classroom.Classroom
var flowCounter int = 0

func main() {
	for flowCounter < 5 {
		switch flowCounter {
		case 0:
			initialization()
			break
		case 1:
			optionAddOrSeeStudent()
			break
		case 2:
			formAddStudent()
			break
		case 3:
			optionSeeStudent()
			break
		case 4:
			optionSeeHighestAndLowest()
			break
		}
	}
	appearance.ShowGoodbyeMessage()
}

func initialization() {
	clsrm = classroom.Classroom{}
	appearance.ShowInitialOption()
	flowCounter += input.InputProgramContinue()
}

func optionAddOrSeeStudent() {
	appearance.ShowOptionAddStudent()
	flowCounter += input.InputContinueFlow(3)
}

func formAddStudent() {
	appearance.ShowAddStudentDetails()
	clsrm.AddClassroomStudents()
	flowCounter -= 1
}

func optionSeeStudent() {
	appearance.ShowStudentsOrderedByGpa(clsrm)
	students := clsrm.ClassroomStudents()
	inputContinue := input.InputContinueFlow(len(students))
	if inputContinue == -1 {
		flowCounter += -2
	} else if inputContinue == 99 {
		flowCounter += 99
	} else {
		fmt.Printf("there are: %d\n", inputContinue-1)
		appearance.ShowStudentDetails(students[inputContinue-1])
		inputContinue = input.InputContinueFlow(1)
		if inputContinue == -1 {
			flowCounter += -2
		} else if inputContinue == 99 {
			flowCounter += 99
		} else {
			flowCounter = 0
		}
	}
}

func optionSeeHighestAndLowest() {
	appearance.ShowHighestAndLowestGpa(clsrm)
	inputContinue := input.InputContinueFlow(len(clsrm.ClassroomStudents()))
	if inputContinue == -1 {
		flowCounter -= 3
	} else {
		flowCounter += inputContinue
	}
}
