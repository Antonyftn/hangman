package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const maxErrors = 6

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <fichier_mots>")
		os.Exit(1)
	}

	words, err := readWordsFile(os.Args[1])
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier: %v\n", err)
		os.Exit(1)
	}

	word := chooseRandomWord(words)
	hiddenWord := initializeHiddenWord(word)
	usedLetters := make(map[rune]bool)
	errors := 0

	hiddenWord = revealRandomLetters(word, hiddenWord, 2) // Révéler 2 lettres aléatoires

	reader := bufio.NewReader(os.Stdin)
	for errors < maxErrors && !isWordComplete(hiddenWord) {
		displayGameState(hiddenWord, usedLetters, errors)

		fmt.Print("Entrez une lettre ou un mot complet: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if len(input) == 1 {
			letter := rune(input[0])
			if !handleDuplicateLetter(letter, usedLetters) {
				if strings.ContainsRune(word, letter) {
					hiddenWord = updateHiddenWord(word, hiddenWord, letter)
				} else {
					errors++
					fmt.Println("Lettre incorrecte!")
				}
				usedLetters[letter] = true
			}
		} else if len(input) > 1 {
			if input == word {
				hiddenWord = word
			} else {
				errors++
				fmt.Println("Mot incorrect!")
			}
		}
	}

	displayEndGame(hiddenWord, word)
}