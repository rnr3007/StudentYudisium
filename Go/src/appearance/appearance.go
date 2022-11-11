package appearance

import (
	"fmt"
	"main/college/classroom"
	"main/user"
	"os"
	"os/exec"
	"runtime"
)

func ShowInitialOption() {
	showAppDescription()
	fmt.Printf("| Do you want to continue? (Y/N)                |\n")
	showBottomBorder()
}

func ShowOptionAddStudent() {
	showAppDescription()
	fmt.Printf("| (1) Add student                               |\n")
	fmt.Printf("| (2) Show student                              |\n")
	fmt.Printf("| (3) Show student with highest and lowest GPA  |\n")
	showAppGeneralMenu()
}

func ShowAddStudentDetails() {
	showAppDescription()
	fmt.Printf("| Please input student name and course amount   |\n")
	fmt.Printf("| Press [ENTER] after you done                  |\n")
	showBottomBorder()
}

func ShowStudentForm(name string, indexCourse int) {
	showAppDescription()
	fmt.Printf("| Hi, %-41s |\n", name)
	fmt.Printf("|                                               |\n")
	fmt.Printf("| Please input course name, sks, and            |\n")
	fmt.Printf("| letter mark for COURSE %-2d                     |\n", indexCourse)
	fmt.Printf("| Press [ENTER] after you done                  |\n")
	showBottomBorder()
}

func ShowStudentsOrderedByGpa(clsrm classroom.Classroom) {
	var students []user.Student
	students = clsrm.ClassroomStudents()
	showAppDescription()
	if len(students) != 0 {
		fmt.Printf("| Students inside this classroom ordered by     |\n")
		fmt.Printf("| their GPA:                                    |\n")
		fmt.Printf("|                                               |\n")
		for index := 0; index < len(students); index++ {
			fmt.Printf("| %2d. %-36s %.2f |\n", index+1,
				students[index].StudentName(), students[index].StudentGpa())
		}
		fmt.Printf("|                                               |\n")
		fmt.Printf("| Choose the student number to look more detail |\n")
		fmt.Printf("| value                                         |\n")
	} else {
		fmt.Printf("| No student found! Please add student first!   |\n")
	}
	showAppGeneralMenu()
}

func ShowStudentDetails(stdnt user.Student) {
	courses := stdnt.StudentCourses()
	counter := 1
	showAppDescription()
	fmt.Printf("| >>Name     : %-32s |\n", stdnt.StudentName())
	fmt.Printf("| >>Predicate: %-32s |\n", stdnt.StudentPredicate())
	fmt.Printf("| >>Courses  :                                  |\n")
	for _, crs := range courses {
		fmt.Printf("|   %2d. %-35s (%s) |\n", counter, crs.CourseName(), crs.CourseLetterMark())
		counter++
	}
	fmt.Printf("| >>Total SKS: %-2d                               |\n", stdnt.StudentSksTotal())
	fmt.Printf("| >>GPA      : %-1.2f                             |\n", stdnt.StudentGpa())
	fmt.Printf("|                                               |\n")
	fmt.Printf("| (1) Restart app                               |\n")
	showAppGeneralMenu()
}

func ShowHighestAndLowestGpa(clsrm classroom.Classroom) {
	var students []user.Student
	students = clsrm.ClassroomStudents()
	showAppDescription()
	if len(students) != 0 {
		fmt.Printf("| Students with the highest GPA is:             |\n")
		fmt.Printf("| %-40s %.2f |\n",
			students[0].StudentName(), students[0].StudentGpa())
		fmt.Printf("|                                               |\n")
		fmt.Printf("| Students with the lowest GPA is:              |\n")
		fmt.Printf("| %-40s %.2f |\n",
			students[len(students)-1].StudentName(), students[len(students)-1].StudentGpa())
	} else {
		fmt.Printf("| No student found! Please add student first!   |\n")
	}
	showAppGeneralMenu()
}

func ShowGoodbyeMessage() {
	showAppDescription()
	fmt.Printf("|          The application is stopped!          |\n")
	fmt.Printf("|                                               |\n")
	fmt.Printf("|     Thank you for using our application!      |\n")
	showBottomBorder()
}

func showAppDescription() {
	clearTerminalDisplay()
	fmt.Printf("=================================================\n")
	fmt.Printf("|                                               |\n")
	fmt.Printf("|    Welcome to the Student GPA and Yudisium    |\n")
	fmt.Printf("|                  checker app                  |\n")
	fmt.Printf("|                                               |\n")
	fmt.Printf("| We'll check your graduation status and count  |\n")
	fmt.Printf("| your GPA                                      |\n")
	fmt.Printf("|                                               |\n")
}

func showAppGeneralMenu() {
	fmt.Printf("|                                               |\n")
	fmt.Printf("| (0) Back                                      |\n")
	fmt.Printf("| (99) Close the app                            |\n")
	showBottomBorder()
}

func showBottomBorder() {
	fmt.Printf("|                                               |\n")
	fmt.Printf("=================================================\n")
}

func clearTerminalDisplay() {

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
