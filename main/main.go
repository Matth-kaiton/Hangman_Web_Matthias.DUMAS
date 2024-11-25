package main

import (
	"hangman"
	"strings"
)

func main() {
	word := strings.ToUpper(hangman.ReadFile(hangman.Diff()))
	wordSlice := hangman.ScliceWord(word)
	hangman.Play(word, wordSlice)

}
