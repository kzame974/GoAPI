package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kzame974/GoAPI/src/models"
	"github.com/kzame974/GoAPI/src/services/influxService"
	"net/http"
)

func PostTemperatureController(c *gin.Context) {

	var temperature models.Temperature
	if err := c.ShouldBindJSON(&temperature); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if influxService.WriteDataToInfluxDB(c, temperature) {
		//reponse ok
		c.JSON(http.StatusOK, gin.H{"message": "Données ajoutées avec succès"})
	}
}
