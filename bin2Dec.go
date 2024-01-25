package main

import (
	"fmt"
	"math"
	"strconv"
	"regexp"
)

func main() {
	// Welcome message
	fmt.Println("Welcome to the binary 2 Decimal converter!")

	// grab user input as int
	var userInput int
	fmt.Print("Enter a binary number: ")
	fmt.Scan(&userInput)

	// convert input to string to check if it contains anything other than 1 or 0
	inputString := strconv.Itoa(userInput)
	notValidInput, _ := regexp.MatchString("[2-9]", inputString)

	if notValidInput {
		fmt.Println("You can only enter combonations of 0 or 1 for a binary num")
	} else {
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
}