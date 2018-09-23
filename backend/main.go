package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Print("Started cms backend 😎")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
