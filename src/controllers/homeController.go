package controllers

import "github.com/gin-gonic/gin"

// Juste pour un petit clin d'oeil ;)
func HomeController(c *gin.Context) {
	c.String(200, "Oté comment y lé Feel bat ?")
}
