package main

import (
    "math/rand"
    "time"
)

func chooseRandomWord(words []string) string {
    rand.Seed(time.Now().UnixNano())
    return words[rand.Intn(len(words))]
}

func revealLetters(word string, numRevealed int) string {
    revealed := make([]rune, len(word))
    for i := range revealed {
        revealed[i] = '_'
    }

    rand.Seed(time.Now().UnixNano())
    revealedIndices := rand.Perm(len(word))[:numRevealed]
    for _, index := range revealedIndices {
        revealed[index] = rune(word[index])
    }

    return string(revealed)
}
