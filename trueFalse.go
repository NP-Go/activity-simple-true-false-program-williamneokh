package main

import "fmt"

func compare(value int) string {
	//do not change this variable resultMessage, secretValue
	resultMessage := ""
	secretValue := 88

	if value == secretValue {
		fmt.Println("Well done! your guess is correct.")
	}
	if value < secretValue {
		fmt.Println("Too low, try again next time!")
	}

	if value > secretValue {
		fmt.Println("Too high, try again next time!")
	}

	//Insert your code from here

	//do not remove this line
	return resultMessage
}

func main() {
	var guess int
	fmt.Println("Enter an integer value from 1 to 100: ")
	fmt.Scanln(&guess)
	compare(guess)
}
