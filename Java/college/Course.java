package college;

import util.InputValidator;

public class Course extends InputValidator{
	private String name;
	private int sks;
	private String letterMark;
	
	public Course(){
		inputName();
		inputSks();
		inputLetterMark();
	}
	
	private void inputName(){
		System.out.printf("\nCourse name: ");
		this.name = safeAlphabetString();
	}
	
	private void inputSks(){
		System.out.printf("\nMany sks (2 - 6): ");
		do {
			this.sks = safeInteger();
			if(sks<2 || sks>6){
				System.out.printf("Input the proper sks amount (2 - 6): ");
			}
		}
		while (sks<2 || sks>6);
	}
	
	private void inputLetterMark(){
		do {
			System.out.printf("\nCourse letter mark (A/B/C/D/E): ");
			this.letterMark = safeAlphabetString().toUpperCase();
			if(!this.letterMark.matches("\\b[ABCDE]{1}\\b")){
				System.out.printf("Input the proper letter mark (A/B/C/D/E)!\n");
			}
		}
		while (!this.letterMark.matches("\\b[ABCDE]{1}\\b"));
	}
	
	public String getName(){
		return this.name;
	}
	
	public int getSks(){
		return this.sks;
	}
	
	public String getLetterMark(){
		return this.letterMark;
	}
	
	public double getDoubleMark(){
		switch (this.letterMark){
			case "A":
				return 4.0;
			case "B":
				return 3.0;
			case "C":
				return 2.0;
			case "D":
				return 1.0;
			case "E":
				return 0.0;
		}
		return 0d;
	}
}