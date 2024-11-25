package hangman

import (
	"io"
	"os"
)

func ReadHangman() []string {
	var liste []string

	f, e := os.Open("hangman.txt")
	if e != nil {
		panic("Erreur avec le ficher hangman")
	}

	defer f.Close()

	tableau, err := io.ReadAll(f)
	if err != nil {
		panic("io a merdé")
	}

	for i := 0; i < len(tableau); i += 71 {
		liste = append(liste, string(tableau[i:i+70]))
	}

	return liste
}
