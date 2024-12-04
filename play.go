package hangman

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type SaveH struct {
	Word       string
	WordToFind []string
	PV         int
	IsStart    bool
	Input      string
	Used       []string
	Message    string
}

func Play(C SaveH) SaveH {
	digit := false
	symbol := false
	isInWord := false
	var lettre []rune
	var use bool
	var win bool

	// var save *SaveH = &SaveH{
	// 	Word:       C.Word,
	// 	WordToFind: C.WordToFind,
	// 	input:      C.Input,
	// 	PV:         C.PV,
	// 	IsStart:    C.IsStart,
	// 	Used:       C.Used,
	// }

	hangMan := ReadHangman()

	digit = false
	symbol = false
	isInWord = false
	use = false

	if C.IsStart {
		for i := 0; i < len(C.WordToFind); i++ {
			fmt.Print(C.WordToFind[i] + " ")
			if C.WordToFind[i] != "_" {
				C.Used = append(C.Used, C.WordToFind[i])
				C.Used = append(C.Used, strings.ToLower(C.WordToFind[i]))
			}
		}
		C.IsStart = false
	}
	C.Message = ""

	if strings.ToUpper(C.Input) == "STOP" {
		for i := 3; i > 0; i-- {
			fmt.Println("Le partie à été sauvegarder vous aller quitter le jeu dans", i, "seconde")
			time.Sleep(1000 * time.Millisecond)
		}
		SaveQuit(C.Word, C.WordToFind, C.PV, C.IsStart, C.Used, "save.txt")

	}
	for _, i := range C.Input {
		if rune(i) >= 123 || unicode.IsPunct(i) {
			symbol = true
		} else if unicode.IsNumber(i) {
			digit = true
		}
	}
	if len(C.Input) == 1 && !digit && !symbol {
		lettre = []rune(C.Input)

		if lettre[0] >= 'a' && lettre[0] <= 'z' {
			lettre[0] -= 32
			for _, i := range C.Used {
				if i == C.Input {
					use = true
					C.Message = "Cette lettre a déja été utilisé"
					return C
				}
			}
			if !use {
				C.Used = append(C.Used, C.Input)
				C.Used = append(C.Used, strings.ToLower(C.Input))

			}

			for i := 0; i < len(C.Word); i++ {
				if string(lettre[0]) == string(C.Word[i]) {
					C.WordToFind[i] = string(C.Word[i])
					isInWord = true
				}
			}

			for i := 0; i < len(C.WordToFind); i++ {
				if C.WordToFind[i] != "_" {
					C.Used = append(C.Used, C.WordToFind[i])
					C.Used = append(C.Used, strings.ToLower(C.WordToFind[i]))
				}
			}

			if !isInWord {
				C.PV--
				C.Message += strings.ToUpper(C.Input)
				C.Message += " n'est pas présent dans le mot, "
				C.Message += strconv.Itoa(C.PV)
				C.Message += " vie restante"
				fmt.Println(hangMan[9-C.PV])
			}
		}

	} else if len(C.Input) > 1 && !digit && !symbol {
		for _, i := range C.Used {
			if i == C.Input {
				use = true
				C.Message = "Ce mot a déja été utilisé"
				return C
			}
		}
		if !use {
			C.Used = append(C.Used, C.Input)
		}

		if strings.ToUpper(C.Input) == C.Word {
			for i := 0; i < len(C.Word); i++ {
				C.WordToFind[i] = string(C.Word[i])
			}
		} else {

			C.PV -= 2
			C.Message += strings.ToUpper(C.Input)
			C.Message += " Ce n'est pas le bon mot, "
			C.Message += strconv.Itoa(C.PV)
			C.Message += " vie restante"
			fmt.Println(hangMan[9-C.PV])
		}

	} else if digit && symbol {
		C.Message = "Ne rentré pas de caractère spéciaux ou chiffre"
	} else if digit {
		C.Message = "Ne rentré pas de chiffre\n"

	} else if symbol {
		C.Message = "Ne rentré pas de caractère spéciaux"

	}

	for i := 0; i < len(C.Word); i++ {
		if C.WordToFind[i] == string(C.Word[i]) {
			if i == len(C.Word)-1 {
				win = true
			}
		} else {
			break
		}
	}
	if win {
		C.Message = "Vous avez gagné bien joué !!!"
		win = false
	} else if C.PV == 0 {
		C.Message = "Vous avez perdu !!!"
	}
	return C
}
