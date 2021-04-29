package main

import (
	"github.com/deepakputhraya/mailcheck"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})
	r.GET("/validate", func(c *gin.Context) {
		email := c.Query("email")
		if len(strings.TrimSpace(email)) == 0 {
			c.JSON(400, gin.H{
				"success": false,
				"message": "email query param is missing!",
			})
			return
		}
		details, err := mailcheck.GetEmailDetails(email)
		if err != nil {
			log.WithError(err).Error("There was an error")
			c.JSON(500, gin.H{
				"success": false,
				"message": "Internal Server Error",
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"details": details,
		})
	})
	r.Run()
}
