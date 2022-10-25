package asset;

import java.util.List;
import java.util.ArrayList;
import college.*;
import user.Student;

public class Appearance{
	
	protected static void showInitialOption(){
		showAppDescription();
		System.out.printf("| Do you want to continue? (Y/N)                |\n");
		showBottomBorder();
	}
	
	protected static void showOptionAddStudent(){
		showAppDescription();
		System.out.printf("| (1) Add student                               |\n");
		System.out.printf("| (2) Show student                              |\n");
		System.out.printf("| (3) Show student with highest and lowest GPA  |\n");
		showAppGeneralMenu();
	}
	
	protected static void showAddStudentDetails(){
		showAppDescription();
		System.out.printf("| Please input student name and course amount   |\n");
		System.out.printf("| Press [ENTER] after you done                  |\n");
		showBottomBorder();
	}
	
	public static void showAddCourseDetails(String name, int indexCourse){
		showAppDescription();
		System.out.printf("| Hi, %-41s |\n", name);
		System.out.printf("|                                               |\n");
		System.out.printf("| Please input course name, sks, and            |\n");
		System.out.printf("| letter mark for COURSE %-2d                     |\n", indexCourse);
		System.out.printf("| Press [ENTER] after you done                  |\n");
		showBottomBorder();
	}
	
	protected static void showStudentForm(String name, int indexCourse){
		showAppDescription();
		System.out.printf("| Hi, %-41s |\n", name);
		System.out.printf("|                                               |\n");
		System.out.printf("| Please input course name, sks, and            |\n");
		System.out.printf("| letter mark for COURSE %-2d                     |\n", indexCourse);
		System.out.printf("| Press [ENTER] after you done                  |\n");
		showBottomBorder();
	}
	
	protected static void showStudentsOrderedByGpa(Classroom classroom){
		List<Student> students = classroom.getStudents();
		showAppDescription();
		if (!students.isEmpty()){
			System.out.printf("| Students inside this classroom ordered by     |\n");
			System.out.printf("| their GPA:                                    |\n");
			System.out.printf("|                                               |\n");
			for (int index=0; index<students.size(); index++){
				System.out.printf("| %2d. %-36s %.2f |\n", index+1, 
					students.get(index).getName(), students.get(index).getGpa());
			}
			System.out.printf("|                                               |\n");
			System.out.printf("| Choose the student number to look more detail |\n");
			System.out.printf("| value                                         |\n");
		}
		else {
			System.out.printf("| No student found! Please add student first!   |\n");
		}
		showAppGeneralMenu();	
	}
	
	protected static void showStudentDetails(Student student){
		List<Course> courses = student.getCourses();
		int counter = 1;
		showAppDescription();
		System.out.printf("| >>Name     : %-32s |\n", student.getName());
		System.out.printf("| >>Predicate: %-32s |\n", student.getPredicate());
		System.out.printf("| >>Courses  :                                  |\n");
		for (Course course : courses){
			System.out.printf("|   %2d. %-35s (%s) |\n", counter, course.getName(), course.getLetterMark());
			counter++;
		}
		System.out.printf("| >>Total SKS: %-2d                               |\n", student.getSksTotal());
		System.out.printf("| >>GPA      : %-1.2f                             |\n", student.getGpa());
		System.out.printf("|                                               |\n");
		System.out.printf("| (1) Restart app                               |\n");
		showAppGeneralMenu();
	}
	
	protected static void showHighestAndLowestGpa(Classroom classroom){
		List<Student> students = classroom.getStudents();
		showAppDescription();
		if (!students.isEmpty()){
			System.out.printf("| Students with the highest GPA is:             |\n");
			System.out.printf("| %-40s %.2f |\n", 
				students.get(0).getName(), students.get(0).getGpa());
			System.out.printf("|                                               |\n");
			System.out.printf("| Students with the lowest GPA is:              |\n");
			System.out.printf("| %-40s %.2f |\n",
				students.get(students.size()-1).getName(), students.get(students.size()-1).getGpa());
		}
		else {
			System.out.printf("| No student found! Please add student first!   |\n");
		}
		showAppGeneralMenu();	
	}
	
	protected static void showGoodbyeMessage(){
		showAppDescription();
		System.out.printf("|          The application is stopped!          |\n");
		System.out.printf("|                                               |\n");
		System.out.printf("|     Thank you for using our application!      |\n");
		showBottomBorder();	
	}
	
	private static void showAppDescription(){
		clearTerminalDisplay();
		System.out.printf("=================================================\n");
		System.out.printf("|                                               |\n");
		System.out.printf("|    Welcome to the Student GPA and Yudisium    |\n");
		System.out.printf("|                  checker app                  |\n");
		System.out.printf("|                                               |\n");
		System.out.printf("| We'll check your graduation status and count  |\n");
		System.out.printf("| your GPA                                      |\n");
		System.out.printf("|                                               |\n");
	}
		
	private static void showAppGeneralMenu(){
		System.out.printf("|                                               |\n");
		System.out.printf("| (0) Back                                      |\n");
		System.out.printf("| (99) Close the app                            |\n");
		showBottomBorder();
	}
	
	private static void showBottomBorder(){
		System.out.printf("|                                               |\n");
		System.out.printf("=================================================\n");
	}
	
	private static void clearTerminalDisplay(){
		try {
			if (System.getProperty("os.name").contains("Windows")){
				new ProcessBuilder("cmd", "/c", "cls").inheritIO().start().waitFor();
			}
			else {
				Runtime.getRuntime().exec("clear");
			}
		}
		catch (Exception exception){
			System.out.printf("Sorry, you're currently using unknown OS!");
		}
	}
}