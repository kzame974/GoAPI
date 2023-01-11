package controllers

import "github.com/gin-gonic/gin"

// on recupère toutes les temperature dans la base de influx et on retourne la donnée au client en JSON
func GetTemperatureController(c *gin.Context) {
	c.String(200, "Hello world !")
}
