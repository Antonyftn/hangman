package hangman

import (
	"fmt"
	"strings"
)

var hangmanArt = []string{
	`
  +---+
  |   |
      |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}

func DisplayHangman(attempts int) {
	if attempts >= 0 && attempts < len(hangmanArt) {
		fmt.Println(hangmanArt[attempts])
	}
}

func DisplayWord(letters []string) {
	word := strings.Join(letters, " ")
	fmt.Println()
	fmt.Println("Mot : ", word)
	fmt.Println(strings.Repeat("-", len(word)+7))
}

func DisplayTitle() {
	title := `
    _    _          _   _  _____ __  __          _   _ 
   | |  | |   /\   | \ | |/ ____|  \/  |   /\   | \ | |
   | |__| |  /  \  |  \| | |  __| \  / |  /  \  |  \| |
   |  __  | / /\ \ | . ' | | |_ | |\/| | / /\ \ | . ' |
   | |  | |/ ____ \| |\  | |__| | |  | |/ ____ \| |\  |
   |_|  |_/_/    \_\_| \_|\_____|_|  |_/_/    \_\_| \_|
	`
	fmt.Println(title)
}

func DisplayGameOver(won bool, word string) {
	if won {
		fmt.Println(`
   ____                            _       _ 
  / ___| __ _  __ _ _ __   ___    | | ___ | |
 | |  _ / _' |/ _' | '_ \ / _ \   | |/ _ \| |
 | |_| | (_| | (_| | | | |  __/   | | (_) |_|
  \____|\__,_|\__, |_| |_|\___|   |_|\___/(_)
              |___/                          
		`)
	} else {
		fmt.Println(`
   ____                         ___                 _ 
  / ___| __ _ _ __ ___   ___   / _ \__   _____ _ __| |
 | |  _ / _' | '_ ' _ \ / _ \ | | | \ \ / / _ \ '__| |
 | |_| | (_| | | | | | |  __/ | |_| |\ V /  __/ |  |_|
  \____|\__,_|_| |_| |_|\___|  \___/  \_/ \___|_|  (_)
		`)
	}
	fmt.Printf("Le mot était : %s\n", word)
}

func DisplayAttempts(attempts, maxAttempts int) {
	fmt.Printf("Essais restants : %d/%d\n", maxAttempts-attempts, maxAttempts)
}

func DisplayGuessedLetters(guessedLetters []string) {
	fmt.Println("Lettres déjà essayées :", strings.Join(guessedLetters, ", "))
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
