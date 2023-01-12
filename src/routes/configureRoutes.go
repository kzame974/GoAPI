package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kzame974/GoAPI/src/controllers"
)

// ConfigureRoutes On regroupe toutes les routes ici et onappelle la func dans le main avant de lancer le serveur web
func ConfigureRoutes(r *gin.Engine) {
	r.POST("/sensors", controllers.PostSensorDataController)
	r.GET("/", controllers.HomeController)
}
