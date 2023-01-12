package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kzame974/GoAPI/src/models"
	"github.com/kzame974/GoAPI/src/services/influxService"
	"net/http"
)

func PostSensorDataController(c *gin.Context) {

	type HexData struct {
		Hex string `json:"hex"`
	}

	// structure pour stocker les données hexadécimales reçues
	var hexData HexData

	// vérification de la présence des données hexadécimales dans le corps de la requête
	if err := c.ShouldBindJSON(&hexData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//on fait ensuite un décodage de l'hexa. Partons du principe que les données seront toujours dans le même ordre
	decodedSensorData, errDecodeHex := models.DecodeHexData(hexData.Hex)
	if errDecodeHex != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDecodeHex.Error()})
		fmt.Println(errDecodeHex.Error())
	}

	if influxService.WriteDataToInfluxDB(c, decodedSensorData) {
		//reponse ok
		c.JSON(http.StatusOK, gin.H{"message": "Données ajoutées avec succès"})
	}
}
