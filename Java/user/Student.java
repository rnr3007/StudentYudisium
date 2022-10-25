package user;

import java.util.List;
import java.util.ArrayList;
import asset.Appearance;
import util.InputValidator;
import college.Course;

public class Student extends InputValidator{
	private String name;
	private int courseAmount;
	private List<Course> courses = new ArrayList<Course>();
	private int sksTotal;
	private double gpa;
	private String predicate;
	
	public Student(){
		inputName();
		inputCourseAmount();
		inputCourses();
	}
	
	public void inputName(){
		System.out.printf("\nYour name: ");
		this.name = super.safeAlphabetString();
	}
	
	public void inputCourseAmount(){
		System.out.printf("\nMany courses do you have: ");
		do {
			this.courseAmount = super.safeInteger();
			if (this.courseAmount < 2 || this.courseAmount > 12){
				System.out.printf("Enter proper course amount (2 - 12): ");
			}
		}
		while (this.courseAmount < 2 || this.courseAmount > 12);
	}
	
	public void inputCourses(){
		for (int counter=0; counter<this.courseAmount; counter++){
			Appearance.showAddCourseDetails(this.name, counter+1);
			Course course = new Course();
			this.sksTotal += course.getSks();
			this.gpa += (course.getDoubleMark() * course.getSks());
			System.out.printf("\n\n         ;;PRESS [ENTER] TO CONTINUE;;");
			scan.nextLine();
			courses.add(course);
		}
		this.gpa = this.gpa / this.sksTotal;
		setPredicate();
	}
	
	public String getName(){
		return this.name;
	}
	
	public List<Course> getCourses(){
		return this.courses;
	}
	
	public int getSksTotal(){
		return this.sksTotal;
	}
	
	public double getGpa(){
		return this.gpa;
	}
	
	public String getPredicate(){
		return this.predicate;
	}
		
	private void setPredicate(){
		for (Course course : this.courses){
			if (course.getLetterMark().equals("E")){
				this.predicate = "PENDING";
			}
		}
		
		if(this.gpa >= 3.5){
			this.predicate = "CUM LAUDE";
		}
		else if(this.gpa >= 2.75 && this.gpa < 3.5){
			this.predicate = "VERY SATISFYING";
		}
		else if(this.gpa >= 2 && this.gpa < 2.75){
			this.predicate = "SATISFYING";
		}
	}
}