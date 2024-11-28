package hangman

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

type SaveH struct {
	Word       string
	WordToFind []string
	PV         int
	IsStart    bool
	Used       []string
	message    string
}

func Play(s string, o []string) {
	digit := false
	symbol := false
	isInWord := false
	var input string
	var lettre []rune
	win := false
	var use bool

	var save *SaveH = &SaveH{
		Word:       s,
		WordToFind: o,
		PV:         10,
		IsStart:    true,
		Used:       make([]string, 0),
	}

	hangMan := ReadHangman()

	for {
		digit = false
		symbol = false
		isInWord = false
		use = false

		if save.IsStart {

			fmt.Println("\nVous pouvez quitter le jeu a tous moment en écrivant STOP")

			fmt.Println("\n")

			fmt.Println("Good Luck, you have", save.PV, "attempts.\n")
			for i := 0; i < len(save.WordToFind); i++ {
				fmt.Print(save.WordToFind[i] + " ")
				if save.WordToFind[i] != "_" {
					save.Used = append(save.Used, save.WordToFind[i])
					save.Used = append(save.Used, strings.ToLower(save.WordToFind[i]))
				}
			}
			ShowAsci(o)
			fmt.Println("\n")
			save.IsStart = false
		}

		fmt.Print("Choose : ")
		fmt.Scan(&input)
		input = strings.ToLower(input)
		if strings.ToUpper(input) == "STOP" {
			for i := 3; i > 0; i-- {
				fmt.Println("Le partie à été sauvegarder vous aller quitter le jeu dans", i, "seconde")
				time.Sleep(1000 * time.Millisecond)
			}
			SaveQuit(save.Word, save.WordToFind, save.PV, save.IsStart, save.Used, "save.txt")
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
				for _, i := range save.Used {
					if i == input {
						use = true
						fmt.Println("Cette lettre a déja été utilisé")
						break
					}
				}
				if !use {
					save.Used = append(save.Used, input)
					save.Used = append(save.Used, strings.ToLower(input))
				} else {
					continue
				}

				for i := 0; i < len(save.Word); i++ {
					if string(lettre[0]) == string(save.Word[i]) {
						save.WordToFind[i] = string(save.Word[i])
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
			for _, i := range save.Used {
				if i == input {
					use = true
					fmt.Println("Ce mot a déja été utilisé")
					break
				}
			}
			if !use {
				save.Used = append(save.Used, input)
			} else {
				continue
			}

			if strings.ToUpper(input) == save.Word {
				for i := 0; i < len(save.Word); i++ {
					save.WordToFind[i] = string(save.Word[i])
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

		for i := 0; i < len(save.Word); i++ {
			if save.WordToFind[i] == string(save.Word[i]) {
				if i == len(save.Word)-1 {
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
