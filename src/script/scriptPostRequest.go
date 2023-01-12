package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

// génération aléatoire d'une chaîne hexadécimale
func randHex(n int) string {
	var letters = []rune("0123456789abcdef")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// script simple pour envoyer des données dans le body d'un requete POST vers l'endpoint de notre API
func main() {
	// génération aléatoire d'une chaîne hexadécimale
	hexData := randHex(32)

	// création d'une struct contenant les données hexadécimales
	data := struct {
		Hex string `json:"hex"`
	}{
		Hex: hexData,
	}

	// conversion de la struct en json
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling json:", err)
		return
	}

	//Envoi de la requête
	resp, err := http.Post("http://localhost:8083/sensors", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}

	//ferme la réponse à la fin de la fonction pour éviter les fuites de mémoires
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	//Affichage du statut de la réponse
	fmt.Println("POST response:", resp.StatusCode)
}
