package main 

word := chooseRandomWord(words)

	hiddenWord := initializeHiddenWord(word)
	usedLetters := make(map[rune]bool)
	errors := 0

	reader := bufio.NewReader(os.Stdin)
	for errors < maxErrors && !isWordComplete(hiddenWord) {
		displayGameState(hiddenWord, usedLetters, errors)

		fmt.Print("Entrez une lettre ou un mot complet: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if len(input) == 1 {
			letter := rune(input[0])
			if usedLetters[letter] {
				fmt.Println("Vous avez déjà proposé cette lettre.")
				continue
			}
			usedLetters[letter] = true

			if strings.ContainsRune(word, letter) {
				hiddenWord = updateHiddenWord(word, hiddenWord, letter)
			} else {
				errors++
				fmt.Println("Lettre incorrecte!")
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
