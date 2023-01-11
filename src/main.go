package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kzame974/GoAPI/src/routes"
	"log"
	"os"
)

func main() {
	// on définit notre routeur dans le main. Ceci est juste un test pour voir si l'API fonctionne
	router := gin.New()
	port := os.Getenv("PORT")

	errProxy := router.SetTrustedProxies([]string{"192.168.1.2"})
	if errProxy != nil {
		log.Fatal(errProxy)
	}
	routes.UserRoute(router)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
