package input

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SafeInteger() int {
	var integerResult int
	for {
		integer, err := strconv.Atoi(InputValue())
		if err == nil {
			integerResult = integer
			break
		}
		fmt.Printf("Enter the correct integer: ")
	}
	return integerResult
}

func SafeAlphabetString() string {
	var regex, err = regexp.Compile(".?[^A-z ]+.?")
	var alphabetString string = ""
	if err == nil {
		for isOk := true; isOk; isOk = (strings.EqualFold(alphabetString, "") || regex.MatchString(alphabetString)) {
			alphabetString = InputValue()
			if strings.EqualFold(alphabetString, "") || regex.MatchString(alphabetString) {
				fmt.Printf("Enter the correct alphabet string: ")
			}
		}
	}
	return alphabetString
}

func InputValue() string {
	var inputResult string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputResult = scanner.Text()
	}
	return inputResult
}
