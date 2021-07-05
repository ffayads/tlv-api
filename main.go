package main

import (
	"os"

	controllers "github.com/api/tlv-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/getTlv", controllers.GetTlv)
	port := os.Getenv("APIPORT")
	if port == "" {
		port = "5000"
	}
	router.Run(":" + port)
}

//Control de acceso HTTP
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Session-Token, session_token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
	}
}
