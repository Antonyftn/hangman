package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var maxAttempts = 6

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func ReadWordFile(filename string) ([]string, error) {
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

func selectRandomWord(words []string) string {
	if len(words) == 0 {
		return ""
	}
	return words[rand.Intn(len(words))]
}

func toMaj(letter rune) rune {
	if letter >= 97 && letter <= 122 {
		letter -= 32
	}
	return letter
}

func isLetterCorrect(lettre string, mot string) bool {
	res := false
	for _, c := range mot {
		if toMaj(c) == toMaj(rune(lettre[0])) {
			res = true
		}
	}

	return res
}

func updateLetters(letter string, letters []string, word string) {
	for i, l := range word {
		if string(l) == letter {
			letters[i] = letter
		}
	}
}

func InitializeGame(wordFilePath string) []string {
	words, err := ReadWordFile(wordFilePath)
	if err != nil {
		fmt.Println(fmt.Errorf("erreur lors de la lecture du fichier de mots : %v", err))
	}
	if len(words) == 0 {
		fmt.Println(fmt.Errorf("le fichier de mots est vide"))
	}
	return words
}

func LancerLeJeu(words []string) {
	word := selectRandomWord(words)
	letters := make([]string, len(word))
	for i := range letters {
		letters[i] = "_"
		letters[0] = string(word[0])
	}
	attempts := 0
	guessedLetters := make([]string, 0)

	for attempts < maxAttempts {
		ClearScreen()
		DisplayTitle()
		DisplayHangman(attempts)
		DisplayWord(letters)
		DisplayAttempts(attempts, maxAttempts)
		DisplayGuessedLetters(guessedLetters)

		var letter string
		fmt.Print("Entrez une lettre : ")
		fmt.Scanln(&letter)
		letter = strings.ToLower(letter)

		if len(letter) != 1 {
			fmt.Println("Veuillez entrer une seule lettre.")
			continue
		}

		if AlreadyAttempted(letter, guessedLetters) {
			fmt.Println("Vous avez déjà essayé cette lettre.")
			continue
		}

		guessedLetters = append(guessedLetters, letter)

		if isLetterCorrect(letter, word) {
			updateLetters(letter, letters, word)
		} else {
			attempts++
		}

		if IsWon(letters) {
			ClearScreen()
			DisplayGameOver(true, word)
			return
		}
	}

	ClearScreen()
	DisplayGameOver(false, word)
}
