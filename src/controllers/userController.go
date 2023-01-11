package controllers

import "github.com/gin-gonic/gin"

// l'action s'effectue dans le controlleur pour afficher ou donner une réponse. A compléter avec un service et un validator de donnée
func UserController(c *gin.Context) {
	c.String(200, "Hello world !")
}
