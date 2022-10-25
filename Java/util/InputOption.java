package util;

import java.util.Scanner;

public class InputOption extends InputValidator{
	private static Scanner scan = new Scanner(System.in);
	
	public int inputProgramContinue(){
		String continueProgram = "";
		do {
			System.out.printf("\nYour answer: ");
			continueProgram = scan.nextLine();
			if (!(continueProgram.equalsIgnoreCase("Y") || continueProgram.equalsIgnoreCase("N"))){
				System.out.printf("Please input the proper answer!\n");
			}
			else {
				if (continueProgram.equalsIgnoreCase("N")){
					return 99;
				}
			}
		}
		while (!(continueProgram.equalsIgnoreCase("Y") || continueProgram.equalsIgnoreCase("N")));
		return 1;
	}
	
	public int inputContinueFlow(int choiceLimit){
		int choice;
		System.out.printf("\nYour number of choice: ");
		do {
			choice = safeInteger();
			if (choice != 99 && (choice < 0 || choice > choiceLimit)){
				System.out.printf("Input the proper number of choice: ");
			}
		}
		while (choice != 99 && (choice < 0 || choice > choiceLimit));
		return inputBackOrStop(choice);
	}
	
	private int inputBackOrStop(int choice){
		switch (choice){
			case 0:
				return -1;
			case 99:
				return 99;
		}
		return choice;
	}
}