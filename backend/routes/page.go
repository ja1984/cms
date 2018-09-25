package routes

import "github.com/gin-gonic/gin"

func CreatePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetPages(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
