package hangman

func IsWon(letters []string) bool {
	for _, l := range letters {
		if l == "_" {
			return false
		}
	}
	return true
}
