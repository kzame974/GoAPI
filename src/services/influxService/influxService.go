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
)

// test envoie de donnée sur influxDB sans méthode post
func WriteDataToInfluxDB(c *gin.Context, sensor models.SensorTemperatureModel) bool {

	//on se connect au cloud de influx
	writeAPI := ConnectToInfluxDB()

	// création dun point avec les champs "temperature" et "humidity" de type float
	point := influxdb2.NewPoint("sensor",
		map[string]string{},
		map[string]interface{}{
			"temperature": sensor.Temperature,
			"humidity":    sensor.Humidity,
		},
		sensor.Timestamp,
	)
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
	url := os.Getenv("INFLUX_URL")
	client := influxdb2.NewClient(url, token)
	org := os.Getenv("INFLUX_ORG")
	bucket := os.Getenv("INFLUX_BUCKET")

	return client.WriteAPIBlocking(org, bucket)
}
