package main

import (
    "fmt"
    "os"
)

func usedLetters() {
	usedLetters := make(map[rune]bool)
}

func checkLetterUsed(usedLetters map[rune]bool, guess rune) bool {
    if usedLetters[guess] {
        fmt.Println("Lettre déjà proposée.")
        return true
    }
    usedLetters[guess] = true
    return false
}
