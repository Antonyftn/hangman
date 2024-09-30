package main

import (
	"math/rand"
)

func revealRandomLetters(word, hiddenWord string, count int) string {
	revealed := []rune(hiddenWord)
	indexes := rand.Perm(len(word))
	
	for _, i := range indexes {
		if count == 0 {
			break
		}
		if revealed[i] == '_' {
			revealed[i] = rune(word[i])
			count--
		}
	}
	
	return string(revealed)
}
