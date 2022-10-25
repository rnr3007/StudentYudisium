import asset.Appearance;
import college.Classroom;
import user.Student;
import util.*;

public class Main extends Appearance{
	private static InputOption input = new InputOption();
	private static Classroom classroom = new Classroom();
	private static int flowCounter = 0;
	
	public static void main(String[] args){
		while (flowCounter < 5){	
			switch (flowCounter){
				case 0:
					initialization();
					break;
				case 1:
					optionAddOrSeeStudent();
					break;
				case 2:
					formAddStudent();
					break;
				case 3:
					optionSeeStudent();
					break;
				case 4:
					optionSeeHighestAndLowest();
					break;
			}
		}
		showGoodbyeMessage();
	}
	
	private static void initialization(){
		classroom = new Classroom();
		showInitialOption();
		flowCounter += input.inputProgramContinue();
	}
	
	private static void optionAddOrSeeStudent(){
		showOptionAddStudent();
		flowCounter += input.inputContinueFlow(3);
	}
	
	private static void formAddStudent(){
		showAddStudentDetails();
		classroom.addStudent();
		flowCounter -= 1;
	}
	
	private static void optionSeeStudent(){
		classroom.sortStudentByGpa();
		showStudentsOrderedByGpa(classroom);
		int inputContinue = input.inputContinueFlow(classroom.getStudents().size());
		if (inputContinue == -1 || inputContinue == 99){
			flowCounter += (inputContinue == -1) ? -2 : 99;
		}
		else {
			showStudentDetails(classroom.getStudents().get(inputContinue-1));
			inputContinue = input.inputContinueFlow(1);
			if (inputContinue == -1 || inputContinue == 99){
				flowCounter += (inputContinue == -1) ? -2 : 99;
			}
			else {
				flowCounter = 0;
			}
		}
	}
	
	private static void optionSeeHighestAndLowest(){
		classroom.sortStudentByGpa();
		showHighestAndLowestGpa(classroom);
		int inputContinue = input.inputContinueFlow(classroom.getStudents().size());
		if (inputContinue == -1){
			flowCounter -= 3;
		}
		else {
			flowCounter += inputContinue;
		}
	}
}