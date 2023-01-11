package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kzame974/GoAPI/src/routes"
	"log"
	"os"
)

func main() {
	//TODO:: ça fonctionne mais faire une refacto après le POST
	//on définit notre routeur dans le main. Ceci est juste un test pour voir si l'API fonctionne
	router := gin.New()
	port := os.Getenv("PORT")

	//vu dans la doc de go, il faut définir son proxy de confiance, ici on est en localhost
	errProxy := router.SetTrustedProxies([]string{"192.168.1.2"})
	if errProxy != nil {
		log.Fatal(errProxy)
	}
	// gestion de mes routes
	routes.ConfigureRoutes(router)
	err := router.Run(":" + port)
	//err := router.Run(":8083")
	if err != nil {
		log.Fatal(err)
	}
}
