package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// create secret number
	secret := getSecretNumber()

	for matching := false; !matching; {
		// Get user input
		guess := getUserInput()
		fmt.Println(secret, guess)

		// Make comparison (secret and guess)
		matching = isMatching(secret, guess)
		fmt.Println("Matching:", matching)

	}

}

func isMatching(secret, guess int) bool {
	if secret > guess {
		fmt.Println("Your guess is too low")
		return false
	} else if guess > secret {
		fmt.Println("Your guess is too high")
		return false
	} else {
		fmt.Println("You guessed correctly!")
		return true
	}

}

func getSecretNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int() % 11
}

func getUserInput() int {
	fmt.Print("Enter a number: ")
	var input int

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Failed to parse your input")
	} else {
		fmt.Println("You guessed:", input)
	}

	return input
}
