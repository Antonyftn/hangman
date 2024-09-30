package main

import (
    "fmt"
    "os"
)

func hasWon(word, revealed string) bool {
    return word == revealed
}

func handleGuess(word string, revealed []rune, guess rune, errors *int, maxErrors int) bool {
    correctGuess := false
    for i, letter := range word {
        if letter == guess {
            revealed[i] = letter
            correctGuess = true
        }
    }

    if !correctGuess {
        *errors++
        fmt.Printf("Lettre incorrecte. Essais restants : %d\n", maxErrors-*errors)
    }

    return correctGuess
}
