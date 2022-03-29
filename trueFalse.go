package main

import "fmt"

func compare(value int) string {
	//do not change this variable resultMessage, secretValue
	resultMessage := ""
	secretValue := 88

	if value == secretValue {
		return "Well Done! Your guess is correct"
	} else if value < secretValue {
		return "Too low, try again next time!"
	} else if value > secretValue {
		return "Too high, try again next time!"
	} else {
		return resultMessage
	}
	//Insert your code from here

	//do not remove this line

}

func main() {
	var guess int
	fmt.Println("Enter an integer value from 1 to 100: ")
	_, _ = fmt.Scanln(&guess)
	resultMessage := compare(guess)
	fmt.Println(resultMessage)
}
