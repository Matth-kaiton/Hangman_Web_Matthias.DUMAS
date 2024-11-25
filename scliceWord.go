package hangman

import "math/rand"

func ScliceWord(s string) []string {
	var motSlice []string
	var count int
	var positionListe []int
	var position int
	for i := 0; i < len(s); i++ {
		motSlice = append(motSlice, "_")
	}
	countReveal := (len(s) / 2) - 1

	for i := 0; i < countReveal; i++ {
		position = rand.Intn(len(s) - 1)
		if motSlice[position] == "_" {
			lettre := s[position]
			count = 0
			for j := 0; j < len(s); j++ {
				if s[j] == lettre {
					count++
				}
			}
			if (count + len(positionListe)) < countReveal {
				for j := 0; j < len(s); j++ {
					if s[j] == lettre {
						positionListe = append(positionListe, j)
						motSlice[j] = string(s[j])
					}
				}
			} else if (count + len(positionListe)) == countReveal {
				for j := 0; j < len(s); j++ {
					if s[j] == lettre {
						positionListe = append(positionListe, j)
						motSlice[j] = string(s[j])
					}
				}
				return motSlice
			} else {
				i--
			}
		} else {
			i--
		}
	}
	return motSlice
}
