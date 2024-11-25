package hangman

import (
	"bufio"
	"math/rand"
	"os"
)

func ReadFile(r int) string {
	var liste string
	var listeMot []string
	var positionMot int

	switch r {
	case 1:
		liste = "words1.txt"
	case 2:
		liste = "words2.txt"
	case 3:
		liste = "words3.txt"

	}
	f, e := os.Open(liste)
	if e != nil {
		panic("oh non le fichier marche pas !")
	}

	defer f.Close()

	tableau := bufio.NewScanner(f)
	for tableau.Scan() {
		listeMot = append(listeMot, tableau.Text())
	}
	positionMot = rand.Intn(len(listeMot))
	return listeMot[positionMot]
}
