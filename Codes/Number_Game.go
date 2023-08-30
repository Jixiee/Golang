//Write a program to guess a random no between 1 to 100 selected by computer
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100) + 1

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have chosen a random number between 1 and 100.")
	fmt.Println("Try to guess the number!")

	var guess int
	numGuesses := 0

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		numGuesses++

		if guess < target {
			fmt.Println("Too low! Try a higher number.")
		} else if guess > target {
			fmt.Println("Too high! Try a lower number.")
		} else {
			fmt.Printf("Congratulations! You've guessed the number %d in %d guesses.\n", target, numGuesses)
			break
		}
	}
}
