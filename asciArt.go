package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func ShowAsci(s []string) {
	var liste []string
	var runes []rune

	f, e := os.Open("standard.txt")
	if e != nil {
		panic("Erreur avec le ficher asci")
	}

	defer f.Close()

	tableau := bufio.NewScanner(f)

	for tableau.Scan() {
		liste = append(liste, tableau.Text())
	}

	for i := rune(0); i < 9; i++ {
		for j := 0; j < len(s); j++ {
			runes = []rune(s[j])
			if s[j] == "_" {
				fmt.Print(liste[((26)*9)+i], " ")
			} else {
				fmt.Print(liste[((runes[0]-65)*9)+i])
			}

		}
		fmt.Println()

	}

}
