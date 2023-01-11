// TODO ENLEVER MAIN SI ÇA FONCTIONNE
package main

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
	"os"
	"time"
)

// test envoie de donnée sur influxDB sans méthode post
func main() {
	// Création de la connexion à séparer ensuite/////
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "https://eu-central-1-1.aws.cloud2.influxdata.com"
	client := influxdb2.NewClient(url, token)
	org := os.Getenv("INFLUX_ORG")
	bucket := "temperature_data"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	// recupération de la date actuelle
	now := time.Now()

	//Envoi de la donné : measurement: room_temp, location : le lieu, value: la température, le temps
	p := influxdb2.NewPoint("room_temperature3",
		map[string]string{"location": "living_room"},
		map[string]interface{}{"value": 40},
		now)
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Mesure de température envoyée!!! ")

	//Fermeture de la connexion
	err = writeAPI.Flush(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
