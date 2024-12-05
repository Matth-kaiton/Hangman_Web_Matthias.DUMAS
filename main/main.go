package main

import (
	"fmt"
	"hangman"
	"html/template"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../template/index/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var i interface{}
	err = tmpl.Execute(w, i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var CurrentGame = hangman.SaveH{
	PV: 10,
}

func Game(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Get("level") != "" {
		CurrentGame.Word = strings.ToUpper(hangman.ReadFile(r.URL.Query().Get("level")))
		CurrentGame.WordToFind = hangman.ScliceWord(CurrentGame.Word)

	}

	r.ParseForm()
	CurrentGame.Used = append(CurrentGame.Used, r.Form.Get("letter"))
	CurrentGame.Input = r.Form.Get("lettre")

	if CurrentGame.Input != "" {
		CurrentGame = hangman.Play(CurrentGame)
	}

	tmpl, err := template.ParseFiles("../template/index/hangman.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, CurrentGame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("../template/css"))))
	http.Handle("/picture/", http.StripPrefix("/picture", http.FileServer(http.Dir("../template/picture"))))

	fmt.Print("Le serveur est lancer avec le port 8080")

	http.HandleFunc("/", index)
	http.HandleFunc("/hangman", Game)

	http.ListenAndServe(":8080", nil)
}
