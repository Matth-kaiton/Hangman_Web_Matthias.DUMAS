package main

import (
	"fmt"
	"hangman"
	"html/template"
	"net/http"
	"strings"
)

var CurrentGame = hangman.SaveH{
	PV: 10,
}

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

	if CurrentGame.PV == 0 {
		CurrentGame.PV = 10
		http.Redirect(w, r, "/finishLose", http.StatusSeeOther)
	} else if CurrentGame.Message == "Vous avez gagné bien joué !!!" {
		http.Redirect(w, r, "/finishWin", http.StatusSeeOther)
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

func FinishLose(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("../template/index/finishLose.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var i interface {
	}
	err = tmpl.Execute(w, i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusAccepted)
	}
}

func FinishWin(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("../template/index/finishWin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var i interface {
	}
	err = tmpl.Execute(w, i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusAccepted)
	}

}

func main() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("../template/css"))))
	http.Handle("/picture/", http.StripPrefix("/picture", http.FileServer(http.Dir("../template/picture"))))

	fmt.Print("Le serveur est lancer avec le port 8080")

	http.HandleFunc("/", index)
	http.HandleFunc("/hangman", Game)
	http.HandleFunc("/finishLose", FinishLose)
	http.HandleFunc("/finishWin", FinishWin)

	http.ListenAndServe(":8080", nil)
}
