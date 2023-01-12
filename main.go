package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kzame974/GoAPI/src/routes"
	"log"
	"os"
)

func main() {
	// charge les variables d'environnement à partir du .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//on définit notre routeur dans le main. Ceci est juste un test pour voir si l'API fonctionne
	router := gin.New()

	//vu dans la doc de go, il faut définir son proxy de confiance, ici on est en localhost
	errProxy := router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXY")})
	if errProxy != nil {
		log.Fatal(errProxy)
	}
	// gestion de mes routes
	routes.ConfigureRoutes(router)
	err = router.Run(":" + os.Getenv("PORT"))
	//err := router.Run(":8083")
	if err != nil {
		log.Fatal(err)
	}
}
