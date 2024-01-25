package main

import (
	"fmt"
	"math"
)

func main() {
	// Welcome message
	fmt.Println("Welcome to the binary 2 Decimal converter!")

	// grab user input as int
	var userInput int
	fmt.Print("Enter a binary number: ")
	fmt.Scan(&userInput)
	userInputCopy := userInput

	index := 0
	decimalNum := 0
	// convert user input into decimal
	for userInput != 0 {
		remainder := userInput % 10
		userInput = userInput / 10
		decimalNum = decimalNum + remainder * int(math.Pow(2, float64(index)))
		index++
	}

	fmt.Printf("Binary num %v converted to decimal %v\n", userInputCopy, decimalNum)
}