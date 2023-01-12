package influxService

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/kzame974/GoAPI/src/models"
	"net/http"
	"os"
	"time"
)

// test envoie de donnée sur influxDB sans méthode post
func WriteDataToInfluxDB(c *gin.Context, temperature models.Temperature) bool {

	writeAPI := ConnectToInfluxDB()
	// recupération de la date actuelle
	now := time.Now()

	//Envoi de la donné : measurement: room_temp, location : le lieu, value: la température, le temps
	//point := influxdb2.NewPoint("room_temperature3",
	//	map[string]string{"location": "living_room"},
	//	map[string]interface{}{"value": 40},
	//	now)

	point := influxdb2.NewPoint("room_temperature42",
		map[string]string{"location": temperature.Location},
		map[string]interface{}{"value": temperature.Value},
		now)
	// Envoi des données à InfluxDB
	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return false
	}

	//Flush pour s'assurer que les données sont bien envoyées
	if err := writeAPI.Flush(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
		return false
	}
	return true
}

func ConnectToInfluxDB() api.WriteAPIBlocking {
	// Création de la connexion à séparer ensuite/////
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "https://eu-central-1-1.aws.cloud2.influxdata.com"
	client := influxdb2.NewClient(url, token)
	org := os.Getenv("dev team")
	bucket := "temperature_data"

	return client.WriteAPIBlocking(org, bucket)
}
