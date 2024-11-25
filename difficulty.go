package hangman

import (
	"fmt"
	"strconv"
)

func Diff() int {
	var o string
	var i int

	fmt.Print("1 : Facile\n")
	fmt.Print("2 : Moyen\n")
	fmt.Print("3 : Difficile\n")
	fmt.Print("Choisi le numéro de ta difficulté : ")
	fmt.Scan(&o)

	i, _ = strconv.Atoi(o)

	if i >= 1 && i <= 3 {
		switch i {
		case 1:
			fmt.Print("Tu as choisi facile\n")
			return i

		case 2:
			fmt.Print("Tu as choisi moyen\n")
			return i

		case 3:
			fmt.Print("Tu as choisi difficile\n")
			return i
		}
	} else {
		fmt.Print("La dificulté n'existe pas ou vous n'avez pas rentré le bon caractère\n")
		return Diff()

	}
	return 0

}