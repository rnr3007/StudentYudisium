package input

import (
	"fmt"
	"strings"
)

func InputProgramContinue() int {
	continueProgram := ""
	for isOk := true; isOk; isOk = !(strings.EqualFold(continueProgram, "Y") || strings.EqualFold(continueProgram, "N")) {
		fmt.Printf("\nYour answer: ")
		continueProgram = InputValue()
		if !(strings.EqualFold(continueProgram, "Y") || strings.EqualFold(continueProgram, "N")) {
			fmt.Printf("Please input the proper answer!\n")
			continue
		}
	}
	if strings.EqualFold(continueProgram, "N") {
		return 99
	}
	return 1
}

func InputContinueFlow(choiceLimit int) int {
	var choice int
	fmt.Printf("\nYour number of choice: ")
	for isOk := true; isOk; isOk = (choice != 99 && (choice < 0 || choice > choiceLimit)) {
		choice = SafeInteger()
		if choice != 99 && (choice < 0 || choice > choiceLimit) {
			fmt.Printf("Input the proper number of choice: ")
		}
	}
	return inputBackOrStop(choice)
}

func inputBackOrStop(choice int) int {
	switch choice {
	case 0:
		return -1
	case 99:
		return 99
	}
	return choice
}
