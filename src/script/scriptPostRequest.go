package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Temperature est la structure qui contient les données de la temperature
type Temperature struct {
	Time     string  `json:"time"`
	Location string  `json:"location"`
	Value    float32 `json:"value"`
}

func main() {
	//Exemple de données
	data := Temperature{
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
	fmt.Println("POST response status:", resp.StatusCode)
}
