//Write a program to guess the fruit the computer has selected in only 5 attempts
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	words := []string{"apple", "banana", "cherry", "grape", "orange"}
	secretWord := words[rand.Intn(len(words))]

	maxTries := 5
	numTries := 0
	guesses := make([]string, 0)

	fmt.Println("Welcome to the Word Guessing Game!")
	fmt.Println("I'm thinking of a fruit name. Can you guess it?")

	for numTries < maxTries {
		fmt.Printf("You have %d tries left.\nEnter your guess: ", maxTries-numTries)
		var guess string
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}

		guess = strings.ToLower(strings.TrimSpace(guess))
		numTries++
		guesses = append(guesses, guess)

		if guess == secretWord {
			fmt.Printf("Congratulations! You guessed the word '%s' correctly in %d tries.\n", secretWord, numTries)
			fmt.Printf("Your guesses: %s\n", strings.Join(guesses, ", "))
			return
		} else {
			fmt.Println("Try again!")
		}
	}

	fmt.Printf("Sorry, you're out of tries. The secret word was '%s'.\n", secretWord)
	fmt.Printf("Your guesses: %s\n", strings.Join(guesses, ", "))
}
