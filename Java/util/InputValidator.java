package util;

import java.util.Scanner;

public class InputValidator{
	protected static Scanner scan = new Scanner(System.in);

	protected static int safeInteger(){
		while (true){
			try {
				return Integer.valueOf(scan.nextLine());
			}
			catch (Exception e){
				System.out.printf("Enter the correct integer: ");
			}
		}
	}
	
	protected static String safeAlphabetString(){
		String alphabetString;
		do {
			alphabetString = scan.nextLine();
			if (alphabetString.equals("") || alphabetString.matches(".?[^A-z ]+.?")){
				System.out.printf("Enter the correct alphabet string: ");
			}
		}
		while (alphabetString.equals("") || alphabetString.matches(".?[^A-z ]+.?"));
		return alphabetString;
	}
}