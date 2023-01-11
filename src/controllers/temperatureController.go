package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/kzame974/GoAPI/src/services/influxService"
	"net/http"
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

	//Envoi de la donné : measurement: room_temp, location : le lieu, value: la température, le temps
	//p := influxdb2.NewPoint("room_temperature3",
	//	map[string]string{"location": "living_room"},
	//	map[string]interface{}{"value": 40},
	//	now)
	//err := writeAPI.WritePoint(context.Background(), p)
	//if err != nil {
	//	log.Fatal(err)
	//}

	point := influxdb2.NewPointWithMeasurement("room_temperature")

	fmt.Println(point)
	fmt.Println("ahhhhhhhhhhhh")
	// Envoi des données à InfluxDB
	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Flush pour s'assurer que les données sont bien envoyées
	if err := writeAPI.Flush(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//reponse ok
	c.JSON(http.StatusOK, gin.H{"message": "Données ajoutées avec succès"})
}
