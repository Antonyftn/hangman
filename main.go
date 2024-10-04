package main

import (
	"fmt"
	hangman "hangman/doc"
	"os"
)

func main() {
	fichier := "mots.txt"
	args := os.Args[1:]
	if len(args) == 1 {
		fichier = args[0]
	}
	if !hangman.FileExists(fichier) {
		fmt.Println("Erreur: Le fichier" + fichier + " n'existe pas.")
		return
	}
	hangman.AfficherMenu(fichier)
}
