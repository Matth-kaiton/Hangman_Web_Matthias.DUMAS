package hangman

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

type saveH struct {
	word       string
	wordToFind []string
	PV         int
	isStart    bool
	used       []string
}

func Play(s string, o []string) {
	digit := false
	symbol := false
	isInWord := false
	var input string
	var lettre []rune
	win := false
	var use bool

	var save *saveH = &saveH{
		word:       s,
		wordToFind: o,
		PV:         10,
		isStart:    true,
		used:       make([]string, 0),
	}

	hangMan := ReadHangman()

	for {
		digit = false
		symbol = false
		isInWord = false
		use = false

		if save.isStart {

			fmt.Println("\nVous pouvez quitter le jeu a tous moment en écrivant STOP")

			fmt.Println("\n")

			fmt.Println("Good Luck, you have", save.PV, "attempts.\n")
			for i := 0; i < len(save.wordToFind); i++ {
				fmt.Print(save.wordToFind[i] + " ")
				if save.wordToFind[i] != "_" {
					save.used = append(save.used, save.wordToFind[i])
					save.used = append(save.used, strings.ToLower(save.wordToFind[i]))
				}
			}
			ShowAsci(o)
			fmt.Println("\n")
			save.isStart = false
		}

		fmt.Print("Choose : ")
		fmt.Scan(&input)
		input = strings.ToLower(input)
		if strings.ToUpper(input) == "STOP" {
			for i := 3; i > 0; i-- {
				fmt.Println("Le partie à été sauvegarder vous aller quitter le jeu dans", i, "seconde")
				time.Sleep(1000 * time.Millisecond)
			}
			SaveQuit(save.word, save.wordToFind, save.PV, save.isStart, save.used, "save.txt")
			break

		}
		for _, i := range input {
			if rune(i) >= 123 || unicode.IsPunct(i) {
				symbol = true
			} else if unicode.IsNumber(i) {
				digit = true
			}
		}
		if len(input) == 1 && !digit && !symbol {
			lettre = []rune(input)

			if lettre[0] >= 'a' && lettre[0] <= 'z' {
				lettre[0] -= 32
				for _, i := range save.used {
					if i == input {
						use = true
						fmt.Println("Cette lettre a déja été utilisé")
						break
					}
				}
				if !use {
					save.used = append(save.used, input)
					save.used = append(save.used, strings.ToLower(input))
				} else {
					continue
				}

				for i := 0; i < len(save.word); i++ {
					if string(lettre[0]) == string(save.word[i]) {
						save.wordToFind[i] = string(save.word[i])
						isInWord = true
					}
				}
				ShowAsci(o)
				fmt.Println("\n")
				if !isInWord {
					save.PV--
					fmt.Println(strings.ToUpper(input), "n'est pas présent dans le mot,", save.PV, "vie restante\n")
					fmt.Println(hangMan[9-save.PV])
				}
			}

		} else if len(input) > 1 && !digit && !symbol {
			for _, i := range save.used {
				if i == input {
					use = true
					fmt.Println("Ce mot a déja été utilisé")
					break
				}
			}
			if !use {
				save.used = append(save.used, input)
			} else {
				continue
			}

			if strings.ToUpper(input) == save.word {
				for i := 0; i < len(save.word); i++ {
					save.wordToFind[i] = string(save.word[i])
				}
				ShowAsci(o)
			} else {

				save.PV -= 2
				fmt.Println("Ce n'est pas le bon mot,", save.PV, "vie restante\n")
				fmt.Println(hangMan[9-save.PV])
			}

		} else if digit && symbol {
			fmt.Println("Ne rentré pas de caractère spéciaux ou chiffre\n")
		} else if digit {
			fmt.Println("Ne rentré pas de chiffre\n")

		} else if symbol {
			fmt.Println("Ne rentré pas de caractère spéciaux\n")

		}

		for i := 0; i < len(save.word); i++ {
			if save.wordToFind[i] == string(save.word[i]) {
				if i == len(save.word)-1 {
					win = true
				}
			} else {
				break
			}
		}
		if win {
			fmt.Println("Vous avez gagné bien joué !!!")
			break
		} else if save.PV == 0 {
			fmt.Println("Vous avez perdu")
			break
		}

	}
}
