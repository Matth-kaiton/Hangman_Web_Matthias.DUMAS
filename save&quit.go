package hangman

import (
	"encoding/json"
	"os"
)

func SaveQuit(w string, wT []string, P int, iS bool, u []string, f string) error {

	result := make([]interface{}, 5)
	result[0], result[1], result[2], result[3], result[4] = w, wT, P, iS, u

	data, err := json.Marshal(result)
	if err != nil {
		panic("erreur sur l'encodage")
	}
	err = os.WriteFile(f, data, 0644)
	if err != nil {
		panic("Le fichier de sauvegarde n'a pas put ètre crée")
	}
	return err

}
