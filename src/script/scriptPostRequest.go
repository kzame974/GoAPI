package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kzame974/GoAPI/src/models"
	"net/http"
)

// script simple pour envoyer des données dans le body d'un requete POST vers l'endpoint de notre API
func main() {
	//Exemple de données
	data := models.Temperature{
		Time:     "2022-01-01T00:00:00Z",
		Location: "Bathroom",
		Value:    23.5,
	}

	//conversion de la struct en json
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling json:", err)
		return
	}

	//Envoi de la requête
	resp, err := http.Post("http://localhost:8083/temperatures", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	//Affichage de la reponse
	fmt.Println("POST response:", resp.StatusCode)
}
