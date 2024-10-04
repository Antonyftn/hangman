package hangman

import (
	"fmt"
	"os"
)

func AfficherMenu(fichierlistemots string) {
	for {
		fmt.Println("=== Menu Principal du Jeu du Pendu ===")
		fmt.Println("1. Jouer")
		fmt.Println("2. Règles du jeu")
		fmt.Println("3. Quitter")
		fmt.Print("Choisissez une option (1-3) : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fmt.Println("\nLancement du jeu...")
			liste := InitializeGame(fichierlistemots)
			LancerLeJeu(liste)
		case 2:
			fmt.Println("\n=== Règles du Jeu du Pendu ===")
			fmt.Println("1. Un mot secret est choisi aléatoirement.")
			fmt.Println("2. Vous devez deviner le mot en proposant des lettres.")
			fmt.Println("3. Vous avez un nombre limité d'essais avant que le pendu ne soit complété.")
			fmt.Println("4. Si vous devinez le mot avant d'épuiser vos essais, vous gagnez !")
			fmt.Println("5. Si le pendu est complété avant que vous ne deviniez le mot, vous perdez.")
		case 3:
			fmt.Println("Merci d'avoir joué. Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("Option invalide. Veuillez choisir 1, 2 ou 3.")
		}
	}
}
