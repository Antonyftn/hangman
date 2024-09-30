package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

func readWordsFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func chooseRandomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func initializeHiddenWord(word string) string {
	hidden := make([]rune, len(word))
	for i := range hidden {
		hidden[i] = '_'
	}
	return string(hidden)
}

func updateHiddenWord(word, hiddenWord string, letter rune) string {
	result := []rune(hiddenWord)
	for i, char := range word {
		if char == letter {
			result[i] = letter
		}
	}
	return string(result)
}

func isWordComplete(hiddenWord string) bool {
	return !strings.ContainsRune(hiddenWord, '_')
}

func displayGameState(hiddenWord string, usedLetters map[rune]bool, errors int) {
	fmt.Printf("\nMot: %s\n", hiddenWord)
	fmt.Printf("Erreurs: %d/%d\n", errors, maxErrors)
	fmt.Print("Lettres utilisées: ")
	for letter := range usedLetters {
		fmt.Printf("%c ", letter)
	}
	fmt.Println()
}