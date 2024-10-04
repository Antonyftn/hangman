package hangman

import (
	"math/rand"
	"time"
)

func RevealRandomLetters(word string) []string {
	rand.Seed(time.Now().UnixNano())
	letters := make([]string, len(word))
	for i := range letters {
		letters[i] = "_"
	}

	numToReveal := len(word)/2 - 1
	for i := 0; i < numToReveal; i++ {
		index := rand.Intn(len(word))
		if letters[index] == "_" {
			letters[index] = string(word[index])
		} else {
			i-- // Try again if this letter was already revealed
		}
	}

	return letters
}
