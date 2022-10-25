package college;

import java.util.List;
import java.util.ArrayList;
import user.Student;

public class Classroom{
	private List<Student> students = new ArrayList<>();
	
	public void addStudent(){
		Student student = new Student();
		students.add(student);
	}
	
	public List<Student> getStudents(){
		return this.students;
	}
	
	public void sortStudentByGpa(){
		Student tempStudent;
		for (int index=this.students.size()-1; index>0; index--){
			for (int index2=0; index2<index; index2++){
				Student currentStudent = this.students.get(index);
				double currentGpa = currentStudent.getGpa();
				Student startStudent = this.students.get(index2);
				double startGpa = startStudent.getGpa();
				if (currentGpa > startGpa){
					tempStudent = currentStudent;
					this.students.set(index, startStudent);
					this.students.set(index2, tempStudent);
				}
			}
		}
	}
}