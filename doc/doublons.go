package hangman

func AlreadyAttempted(letter string, letters []string) bool {
	for _, l := range letters {
		if l == letter {
			return true
		}
	}
	return false
}
