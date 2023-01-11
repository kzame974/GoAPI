package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kzame974/GoAPI/src/controllers"
)

// le chemin d'accès qu'on a définis pour accéder aux controlleurs
func GetTemperatureRoute(router *gin.Engine) {
	router.GET("/temperature", controllers.GetTemperatureController)
}
