package controllers

import (
	"context"
	"github.com/kzame974/GoAPI/src/services/influxService"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb-client-go/v2"
)

type Temperature struct {
	Location string  `json:"location"`
	Value    float64 `json:"value"`
	Time     string  `json:"time"`
}

func PostTemperatureController(c *gin.Context) {

	var temperature Temperature
	if err := c.ShouldBindJSON(&temperature); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// creation de la connexion
	writeAPI := influxService.ConnectToInfluxDB()

	// conversion de la donnée
	t, err := time.Parse("2006-01-02T15:04:05Z07:00", temperature.Time)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Envoi de la donné : measurement: room_temp, location : le lieu, value: la température, le temps
	//TODO---- regarder la doc pour voir quelle des deux function de poitn est plus pertinente + gérer souci encodage
	point := influxdb2.NewPoint("room_temperature4",
		map[string]string{"location": "bathroom"},
		map[string]interface{}{"value": 10},
		t)
	err = writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		log.Fatal(err)
	}

	//point := influxdb2.NewPointWithMeasurement("room_temperature4")
	//point.AddTag("location", temperature.Location)
	//point.AddField("value", temperature.Value)
	//point.SetTime(t)

	// Envoi des données à InfluxDB
	if err = writeAPI.WritePoint(context.Background(), point); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Flush pour s'assurer que les données sont bien envoyées
	if err = writeAPI.Flush(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//reponse ok
	c.JSON(http.StatusOK, gin.H{"message": "Données ajoutées avec succès"})
}
